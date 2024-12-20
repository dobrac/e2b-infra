package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/e2b-dev/infra/packages/orchestrator/internal/consul"
	"github.com/e2b-dev/infra/packages/orchestrator/internal/dns"
	"github.com/e2b-dev/infra/packages/orchestrator/internal/sandbox"
	"github.com/e2b-dev/infra/packages/shared/pkg/grpc/orchestrator"
	"github.com/e2b-dev/infra/packages/shared/pkg/logs"
	consulapi "github.com/hashicorp/consul/api"

	"go.opentelemetry.io/otel"
)

func main() {
	templateId := flag.String("template", "", "template id")
	buildId := flag.String("build", "", "build id")
	sandboxId := flag.String("sandbox", "", "sandbox id")
	keepAlive := flag.Int("alive", 0, "keep alive")
	count := flag.Int("count", 1, "number of serially spawned sandboxes")

	flag.Parse()

	timeout := time.Second*time.Duration(*keepAlive) + time.Second*50
	fmt.Printf("timeout: %d\n", timeout)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// Start of mock build for testing
	dns := dns.New()
	go dns.Start("127.0.0.4:53")

	consulClient, err := consul.New(context.Background())
	if err != nil {
		panic(err)
	}

	networkPool := sandbox.NewNetworkSlotPool(10, 0)

	go func() {
		poolErr := networkPool.Start(ctx, consulClient)
		if poolErr != nil {
			fmt.Fprintf(os.Stderr, "network pool error: %v\n", poolErr)
		}

		closeErr := networkPool.Close(consulClient)
		if closeErr != nil {
			fmt.Fprintf(os.Stderr, "network pool close error: %v\n", closeErr)
		}
	}()

	eg, ctx := errgroup.WithContext(ctx)

	for i := 0; i < *count; i++ {
		fmt.Printf("Starting sandbox %d\n", i)

		eg.Go(func() error {
			mockSandbox(
				ctx,
				*templateId,
				*buildId,
				*sandboxId+"-"+strconv.Itoa(i),
				dns,
				time.Duration(*keepAlive)*time.Second,
				networkPool,
				consulClient,
			)

			return nil
		})

	}

	err = eg.Wait()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to start sandboxes: %v\n", err)
	}
}

func mockSandbox(
	ctx context.Context,
	templateId,
	buildId,
	sandboxId string,
	dns *dns.DNS,
	keepAlive time.Duration,
	networkPool *sandbox.NetworkSlotPool,
	consulClient *consulapi.Client,
) {
	tracer := otel.Tracer(fmt.Sprintf("sandbox-%s", sandboxId))
	childCtx, _ := tracer.Start(ctx, "mock-sandbox")

	start := time.Now()
	logger := logs.NewSandboxLogger(sandboxId, templateId, "test-team", 2, 512, false)

	sbx, err := sandbox.NewSandbox(
		childCtx,
		tracer,
		consulClient,
		dns,
		networkPool,
		&orchestrator.SandboxConfig{
			TemplateID:         templateId,
			FirecrackerVersion: "v1.7.0-dev_8bb88311",
			KernelVersion:      "vmlinux-5.10.186",
			TeamID:             "test-team",
			BuildID:            buildId,
			HugePages:          true,
			MaxInstanceLength:  1,
			SandboxID:          sandboxId,
		},
		"trace-test-1",
		time.Now(),
		time.Now(),
		logger,
	)
	if err != nil {
		panic(err)
	}

	duration := time.Since(start)

	fmt.Printf("[Sandbox is running] - started in %dms (without network)\n", duration.Milliseconds())

	time.Sleep(keepAlive)

	defer sbx.CleanupAfterFCStop(childCtx, tracer, consulClient, dns, sandboxId)

	sbx.Stop(childCtx, tracer)
}

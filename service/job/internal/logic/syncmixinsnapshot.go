package logic

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/hibiken/asynq"
	"github.com/lixvyang/betxin-micro/service/job/internal/svc"
	"github.com/shopspring/decimal"
	"github.com/zeromicro/go-zero/core/logx"
)

// SyncMixinSnapshotHandler
type SyncMixinSnapshotHandler struct {
	svcCtx *svc.ServiceContext
}

func NewSyncMixinSnapshotHandler(svcCtx *svc.ServiceContext) *SyncMixinSnapshotHandler {
	return &SyncMixinSnapshotHandler{
		svcCtx: svcCtx,
	}
}

// every one minute exec : if return err != nil , asynq will retry
func (l *SyncMixinSnapshotHandler) ProcessTask(ctx context.Context, _ *asynq.Task) error {
	fmt.Printf("shcedule job demo -----> every one second exec \n")
	createdAt, err := getTopSnapshotCreatedAt(l.svcCtx.MixinClient, ctx)
	if err != nil {
		return err
	}
	stats := &Stats{createdAt}
	l.sendTopCreatedAtToChannel(ctx, stats)
	return nil
}

type Memo struct {
	Tid    string `json:"tid"`
	Action int    `json:"action" comment:"0yes 1no"`
}

type Stats struct {
	preCreatedAt time.Time
}

func (s *Stats) getPrevSnapshotCreatedAt() time.Time {
	return s.preCreatedAt
}

func (s *Stats) updatePrevSnapshotCreatedAt(time time.Time) {
	s.preCreatedAt = time
}

func getTopSnapshotCreatedAt(client *mixin.Client, c context.Context) (time.Time, error) {
	snapshots, err := client.ReadSnapshots(c, "", time.Now(), "", 50)
	fmt.Println(len(snapshots))
	if err != nil {
		return time.Now(), err
	}
	if len(snapshots) == 0 {
		return time.Now(), nil
	}
	return snapshots[0].CreatedAt, nil
}

func getTopHundredCreated(client *mixin.Client, c context.Context) ([]*mixin.Snapshot, error) {
	snapshots, err := client.ReadSnapshots(c, "", time.Now(), "", 50)
	if err != nil {
		return nil, err
	}
	return snapshots, nil
}

func (l *SyncMixinSnapshotHandler) sendTopCreatedAtToChannel(ctx context.Context, stats *Stats) {
	preCreatedAt := stats.getPrevSnapshotCreatedAt()
	snapshots, err := getTopHundredCreated(l.svcCtx.MixinClient, ctx)
	if err != nil {
		log.Printf("getTopHundredCreated error")
		return
	}
	var wg sync.WaitGroup
	for _, snapshot := range snapshots {
		if snapshot.CreatedAt.After(preCreatedAt) {
			stats.updatePrevSnapshotCreatedAt(snapshot.CreatedAt)
			if snapshot.Amount.Cmp(decimal.NewFromInt(0)) == 1 && snapshot.Type == "transfer" {
				wg.Add(1)
				go func(s *mixin.Snapshot) {
					defer wg.Done()
					_ = l.handlerNewMixinSnapshot(ctx, s)
				}(snapshot)
			}
		}
	}
	wg.Wait()
}

func (l *SyncMixinSnapshotHandler) handlerNewMixinSnapshot(ctx context.Context, snapshot *mixin.Snapshot) error {
	fmt.Println("同步账单")
	if snapshot.Memo == "" {
		logx.Infow("snapshot.Memo == \"\": ", logx.LogField{Key: "Error: ", Value: "Handle Mixin snapshot error!"})
		return nil
	}

	return nil
}

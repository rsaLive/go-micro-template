package task

import (
	"context"
	"time"

	"go-micro-template/internal/biz"
)

// ApprovalChecker 审批状态检查器
type ApprovalChecker struct {
	castUseCase *biz.CastUseCase
	workRepo    *biz.WorkRepo
}

// NewApprovalChecker 创建审批状态检查器
func NewApprovalChecker(castUseCase *biz.CastUseCase, workRepo *biz.WorkRepo) *ApprovalChecker {
	return &ApprovalChecker{
		castUseCase: castUseCase,
		workRepo:    workRepo,
	}
}

// Start 启动定时检查
func (c *ApprovalChecker) Start(ctx context.Context) {
	ticker := time.NewTicker(15 * time.Minute) // 每15分钟检查一次
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			c.checkPendingApprovals()
		}
	}
}

// checkPendingApprovals 检查待处理的审批
func (c *ApprovalChecker) checkPendingApprovals() {
	// 查询所有状态为"审核中"的作品
	/*  works, err := c.workRepo.GetWorksByStatus(model.WorkStatusReviewing)
	    if err != nil {
	        zap.L().Error("Failed to get works with reviewing status", zap.Error(err))
	        return
	    }

	    for _, work := range works {
	        if work.ApprovalID == "" {
	            continue
	        }

	        // 查询审批状态并更新
	        err := c.castUseCase.QueryApprovalStatusAndUpdate(work.Uuid, work.ApprovalID)
	        if err != nil {
	            zap.L().Error("Failed to check approval status",
	                zap.String("workUuid", work.Uuid),
	                zap.String("approvalID", work.ApprovalID),
	                zap.Error(err))
	        }

	        // 避免请求过于频繁
	        time.Sleep(1 * time.Second)
	    }*/
}

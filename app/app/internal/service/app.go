package service

import (
	"context"
	v1 "dhb/app/app/api"
	"dhb/app/app/internal/biz"
	"dhb/app/app/internal/conf"
	"dhb/app/app/internal/pkg/middleware/auth"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"time"
)

// AppService service.
type AppService struct {
	v1.UnimplementedAppServer

	uuc *biz.UserUseCase
	ruc *biz.RecordUseCase
	log *log.Helper
	ca  *conf.Auth
}

// NewAppService new a service.
func NewAppService(uuc *biz.UserUseCase, ruc *biz.RecordUseCase, logger log.Logger, ca *conf.Auth) *AppService {
	return &AppService{uuc: uuc, ruc: ruc, log: log.NewHelper(logger), ca: ca}
}

// EthAuthorize ethAuthorize.
func (a *AppService) EthAuthorize(ctx context.Context, req *v1.EthAuthorizeRequest) (*v1.EthAuthorizeReply, error) {
	// TODO 有效的参数验证
	userAddress := req.SendBody.Address // 以太坊账户
	if "" == userAddress || 20 > len(userAddress) {
		return nil, errors.New(500, "AUTHORIZE_ERROR", "账户地址参数错误")
	}

	// TODO 验证签名

	// 根据地址查询用户，不存在时则创建
	user, err := a.uuc.GetExistUserByAddressOrCreate(ctx, &biz.User{
		Address: userAddress,
	}, req)
	if err != nil {
		return nil, err
	}

	claims := auth.CustomClaims{
		UserId:   user.ID,
		UserType: "user",
		StandardClaims: jwt2.StandardClaims{
			NotBefore: time.Now().Unix(),              // 签名的生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*7, // 7天过期
			Issuer:    "DHB",
		},
	}
	token, err := auth.CreateToken(claims, a.ca.JwtKey)
	if err != nil {
		return nil, errors.New(500, "AUTHORIZE_ERROR", "生成token失败")
	}

	userInfoRsp := v1.EthAuthorizeReply{
		Token: token,
	}
	return &userInfoRsp, nil
}

// Deposit deposit.
func (a *AppService) Deposit(ctx context.Context, req *v1.DepositRequest) (*v1.DepositReply, error) {
	return &v1.DepositReply{}, nil
}

// UserInfo userInfo.
func (a *AppService) UserInfo(ctx context.Context, req *v1.UserInfoRequest) (*v1.UserInfoReply, error) {
	// 在上下文 context 中取出 claims 对象
	var userId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		userId = int64(c["UserId"].(float64))
	}

	return a.uuc.UserInfo(ctx, &biz.User{
		ID: userId,
	})
}

// RecommendUpdate recommendUpdate.
func (a *AppService) RecommendUpdate(ctx context.Context, req *v1.RecommendUpdateRequest) (*v1.RecommendUpdateReply, error) {
	// 在上下文 context 中取出 claims 对象
	var userId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		userId = int64(c["UserId"].(float64))
	}

	return a.uuc.UpdateUserRecommend(ctx, &biz.User{
		ID: userId,
	}, req)
}

// RewardList rewardList.
func (a *AppService) RewardList(ctx context.Context, req *v1.RewardListRequest) (*v1.RewardListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var userId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		userId = int64(c["UserId"].(float64))
	}

	return a.uuc.RewardList(ctx, req, &biz.User{
		ID: userId,
	})
}

func (a *AppService) RecommendRewardList(ctx context.Context, req *v1.RecommendRewardListRequest) (*v1.RecommendRewardListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var userId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		userId = int64(c["UserId"].(float64))
	}

	return a.uuc.RecommendRewardList(ctx, &biz.User{
		ID: userId,
	})
}

func (a *AppService) FeeRewardList(ctx context.Context, req *v1.FeeRewardListRequest) (*v1.FeeRewardListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var userId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		userId = int64(c["UserId"].(float64))
	}

	return a.uuc.FeeRewardList(ctx, &biz.User{
		ID: userId,
	})
}

func (a *AppService) WithdrawList(ctx context.Context, req *v1.WithdrawListRequest) (*v1.WithdrawListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var userId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		userId = int64(c["UserId"].(float64))
	}

	return a.uuc.WithdrawList(ctx, &biz.User{
		ID: userId,
	}, req.Type)
}

// Withdraw withdraw.
func (a *AppService) Withdraw(ctx context.Context, req *v1.WithdrawRequest) (*v1.WithdrawReply, error) {
	// 在上下文 context 中取出 claims 对象
	var userId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		userId = int64(c["UserId"].(float64))
	}

	return a.uuc.Withdraw(ctx, req, &biz.User{
		ID: userId,
	})
}

// SetBalanceReward .
func (a *AppService) SetBalanceReward(ctx context.Context, req *v1.SetBalanceRewardRequest) (*v1.SetBalanceRewardReply, error) {
	// 在上下文 context 中取出 claims 对象
	var userId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		userId = int64(c["UserId"].(float64))
	}

	return a.uuc.SetBalanceReward(ctx, req, &biz.User{
		ID: userId,
	})
}

// DeleteBalanceReward .
func (a *AppService) DeleteBalanceReward(ctx context.Context, req *v1.DeleteBalanceRewardRequest) (*v1.DeleteBalanceRewardReply, error) {
	// 在上下文 context 中取出 claims 对象
	var userId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		userId = int64(c["UserId"].(float64))
	}

	return a.uuc.DeleteBalanceReward(ctx, req, &biz.User{
		ID: userId,
	})
}

func (a *AppService) AdminRewardList(ctx context.Context, req *v1.AdminRewardListRequest) (*v1.AdminRewardListReply, error) {
	return a.uuc.AdminRewardList(ctx, req)
}

func (a *AppService) AdminUserList(ctx context.Context, req *v1.AdminUserListRequest) (*v1.AdminUserListReply, error) {
	return a.uuc.AdminUserList(ctx, req)
}

func (a *AppService) AdminLocationList(ctx context.Context, req *v1.AdminLocationListRequest) (*v1.AdminLocationListReply, error) {
	return a.uuc.AdminLocationList(ctx, req)
}

func (a *AppService) AdminWithdrawList(ctx context.Context, req *v1.AdminWithdrawListRequest) (*v1.AdminWithdrawListReply, error) {
	return a.uuc.AdminWithdrawList(ctx, req)
}

func (a *AppService) AdminWithdraw(ctx context.Context, req *v1.AdminWithdrawRequest) (*v1.AdminWithdrawReply, error) {
	return a.uuc.AdminWithdraw(ctx, req)
}

func (a *AppService) AdminFee(ctx context.Context, req *v1.AdminFeeRequest) (*v1.AdminFeeReply, error) {
	return a.uuc.AdminFee(ctx, req)
}

func (a *AppService) AdminAll(ctx context.Context, req *v1.AdminAllRequest) (*v1.AdminAllReply, error) {
	return a.uuc.AdminAll(ctx, req)
}

func (a *AppService) AdminUserRecommend(ctx context.Context, req *v1.AdminUserRecommendRequest) (*v1.AdminUserRecommendReply, error) {
	return a.uuc.AdminRecommendList(ctx, req)
}

func (a *AppService) AdminMonthRecommend(ctx context.Context, req *v1.AdminMonthRecommendRequest) (*v1.AdminMonthRecommendReply, error) {
	return a.uuc.AdminMonthRecommend(ctx, req)
}

func (a *AppService) AdminConfig(ctx context.Context, req *v1.AdminConfigRequest) (*v1.AdminConfigReply, error) {
	return a.uuc.AdminConfig(ctx, req)
}

func (a *AppService) AdminConfigUpdate(ctx context.Context, req *v1.AdminConfigUpdateRequest) (*v1.AdminConfigUpdateReply, error) {
	return a.uuc.AdminConfigUpdate(ctx, req)
}

func (a *AppService) AdminWithdrawEth(ctx context.Context, req *v1.AdminWithdrawEthRequest) (*v1.AdminWithdrawEthReply, error) {
	return &v1.AdminWithdrawEthReply{}, nil
}

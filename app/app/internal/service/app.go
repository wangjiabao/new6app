package service

import (
	"context"
	"crypto/md5"
	v1 "dhb/app/app/api"
	"dhb/app/app/internal/biz"
	"dhb/app/app/internal/conf"
	"dhb/app/app/internal/pkg/middleware/auth"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"io"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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
	if "" == userAddress || 20 > len(userAddress) ||
		strings.EqualFold("0x000000000000000000000000000000000000dead", userAddress) {
		return nil, errors.New(500, "AUTHORIZE_ERROR", "账户地址参数错误")
	}

	if "" == req.SendBody.Password || 6 > len(req.SendBody.Password) {
		return nil, errors.New(500, "AUTHORIZE_ERROR", "账户密码必须大于6位")
	}
	// TODO 验证签名
	password := fmt.Sprintf("%x", md5.Sum([]byte(req.SendBody.Password)))

	// 根据地址查询用户，不存在时则创建
	user, err := a.uuc.GetExistUserByAddressOrCreate(ctx, &biz.User{
		Address:  userAddress,
		Password: password,
	}, req)
	if err != nil {
		return nil, err
	}

	claims := auth.CustomClaims{
		UserId:   user.ID,
		UserType: "user",
		Password: password,
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

func (a *AppService) TradeList(ctx context.Context, req *v1.TradeListRequest) (*v1.TradeListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var userId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		userId = int64(c["UserId"].(float64))
	}

	return a.uuc.TradeList(ctx, &biz.User{
		ID: userId,
	})
}

func (a *AppService) TranList(ctx context.Context, req *v1.TranListRequest) (*v1.TranListReply, error) {
	// 在上下文 context 中取出 claims 对象
	var userId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		userId = int64(c["UserId"].(float64))
	}

	return a.uuc.TranList(ctx, &biz.User{
		ID: userId,
	}, req.Type, req.Tran)
}

// PasswordChange withdraw.
func (a *AppService) PasswordChange(ctx context.Context, req *v1.PasswordChangeRequest) (*v1.PasswordChangeReply, error) {
	// 在上下文 context 中取出 claims 对象
	return &v1.PasswordChangeReply{
		Password: fmt.Sprintf("%x", md5.Sum([]byte(req.SendBody.Password))),
	}, nil
}

// Withdraw withdraw.
func (a *AppService) Withdraw(ctx context.Context, req *v1.WithdrawRequest) (*v1.WithdrawReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		userId        int64
		tokenPassword string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		if c["Password"] == nil {
			return nil, errors.New(403, "ERROR_TOKEN", "无效TOKEN")
		}
		userId = int64(c["UserId"].(float64))
		tokenPassword = c["Password"].(string)
	}

	if "" == req.SendBody.Password || 6 > len(req.SendBody.Password) {
		return nil, errors.New(500, "AUTHORIZE_ERROR", "账户密码必须大于6位")
	}
	// TODO 验证签名
	password := fmt.Sprintf("%x", md5.Sum([]byte(req.SendBody.Password)))

	return a.uuc.Withdraw(ctx, req, &biz.User{
		ID:       userId,
		Password: tokenPassword,
	}, password)
}

// Tran tran .
func (a *AppService) Tran(ctx context.Context, req *v1.TranRequest) (*v1.TranReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		userId        int64
		tokenPassword string
	)
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		if c["Password"] == nil {
			return nil, errors.New(403, "ERROR_TOKEN", "无效TOKEN")
		}

		userId = int64(c["UserId"].(float64))
		tokenPassword = c["Password"].(string)
	}

	if "" == req.SendBody.Password || 6 > len(req.SendBody.Password) {
		return nil, errors.New(500, "AUTHORIZE_ERROR", "账户密码必须大于6位")
	}
	// TODO 验证签名
	password := fmt.Sprintf("%x", md5.Sum([]byte(req.SendBody.Password)))

	return a.uuc.Tran(ctx, req, &biz.User{
		ID:       userId,
		Password: tokenPassword,
	}, password)
}

func (a *AppService) GetTrade(ctx context.Context, req *v1.GetTradeRequest) (*v1.GetTradeReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		amountB        int64
		tmpValue       int64
		hbs            float64
		amountFloatHbs float64
		amountFloatCsd float64
		csd            string
		err            error
	)

	amountFloat, _ := strconv.ParseFloat(req.SendBody.Amount, 10)
	amountFloatCsd = amountFloat * 10000000000
	amount, _ := strconv.ParseInt(strconv.FormatFloat(amountFloatCsd, 'f', -1, 64), 10, 64)
	if 10000000000 > amount {
		return nil, errors.New(500, "ERROR_TOKEN", "输入错误")
	}

	csd, err = GetAmountOut(req.SendBody.Amount + "000000000000000000")
	if nil != err {
		fmt.Println(2)
		return nil, errors.New(500, "ERROR_TOKEN", "查询币价错误")
	}
	lenValue := len(csd)
	if 10 > lenValue {
		return nil, errors.New(500, "ERROR_TOKEN", "币价过低")
	}
	tmpValue, _ = strconv.ParseInt(csd[0:lenValue-8], 10, 64)
	if 0 == tmpValue {
		return nil, errors.New(500, "ERROR_TOKEN", "币价过低")
	}

	hbs, err = requestHbsResult()
	if nil != err {
		fmt.Println(1)
		return nil, errors.New(500, "ERROR_TOKEN", "查询币价错误")
	}
	amountFloatHbs = amountFloat * 10
	amountB = int64(amountFloatHbs / hbs * 10000000000)
	if 0 >= amountB {
		return nil, errors.New(500, "ERROR_TOKEN", "币价错误")
	}

	return &v1.GetTradeReply{
		AmountCsd: fmt.Sprintf("%.4f", float64(tmpValue)/float64(10000000000)),
		AmountHbs: fmt.Sprintf("%.4f", float64(amountB)/float64(10000000000)),
	}, nil
}

func (a *AppService) Trade(ctx context.Context, req *v1.WithdrawRequest) (*v1.WithdrawReply, error) {
	// 在上下文 context 中取出 claims 对象
	var (
		userId         int64
		tokenPassword  string
		amountB        int64
		tmpValue       int64
		tmpValue2      int64
		hbs            float64
		amountFloatHbs float64
		amountFloatCsd float64
		csd            string
		err            error
	)

	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		if c["Password"] == nil {
			return nil, errors.New(403, "ERROR_TOKEN", "无效TOKEN")
		}

		userId = int64(c["UserId"].(float64))
		tokenPassword = c["Password"].(string)
	}

	if "" == req.SendBody.Password || 6 > len(req.SendBody.Password) {
		return nil, errors.New(500, "AUTHORIZE_ERROR", "账户密码必须大于6位")
	}
	// TODO 验证签名
	password := fmt.Sprintf("%x", md5.Sum([]byte(req.SendBody.Password)))

	amountFloat, _ := strconv.ParseFloat(req.SendBody.Amount, 10)
	amountFloatCsd = amountFloat * 10000000000
	amount, _ := strconv.ParseInt(strconv.FormatFloat(amountFloatCsd, 'f', -1, 64), 10, 64)
	if 10000000000 > amount {
		return &v1.WithdrawReply{
			Status: "fail",
		}, nil
	}

	//if 0 != amount%10 {
	//	return nil, errors.New(500, "ERROR_TOKEN", "10的整数倍")
	//}

	csd, err = GetAmountOut(req.SendBody.Amount + "000000000000000000")
	if nil != err {
		return nil, errors.New(500, "ERROR_TOKEN", "查询币价错误")
	}
	lenValue := len(csd)
	if 10 > lenValue {
		return nil, errors.New(500, "ERROR_TOKEN", "币价过低")
	}
	tmpValue, _ = strconv.ParseInt(csd[0:lenValue-8], 10, 64)
	if 0 == tmpValue {
		return nil, errors.New(500, "ERROR_TOKEN", "币价过低")
	}

	hbs, err = requestHbsResult()
	if nil != err {
		return nil, errors.New(500, "ERROR_TOKEN", "查询币价错误")
	}
	amountFloatHbs = amountFloat * 10
	amountB = int64(amountFloatHbs / hbs * 10000000000)
	if 0 >= amountB {
		return nil, errors.New(500, "ERROR_TOKEN", "币价错误")
	}

	//csdTrade, err = GetAmountOut(strconv.FormatInt((amount+amount*10)/10000000000, 10) + "000000000000000000")
	//if nil != err {
	//	return nil, errors.New(500, "ERROR_TOKEN", "查询币价错误")
	//}
	//lenCsdTradeValue := len(csdTrade)
	//if 10 > lenCsdTradeValue {
	//	return nil, errors.New(500, "ERROR_TOKEN", "币价过低")
	//}
	//tmpValue2, _ = strconv.ParseInt(csdTrade[0:lenCsdTradeValue-8], 10, 64)
	//if 0 == tmpValue2 {
	//	return nil, errors.New(500, "ERROR_TOKEN", "币价过低")
	//}

	return a.uuc.Trade(ctx, req, &biz.User{
		ID:       userId,
		Password: tokenPassword,
	}, tmpValue, amountB, tmpValue2, password)
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

func GetAmountOut(strAmount string) (string, error) {

	var balString string
	url1 := "https://bnb-bscnews.rpc.blxrbdn.com/"

	for i := 4; i < 16; i++ {
		client, err := ethclient.Dial(url1)
		if err != nil {
			return "", err
		}

		tokenAddress := common.HexToAddress("0x10ED43C718714eb63d5aA57B78B54704E256024E")
		instance, err := NewPancakerouterv2(tokenAddress, client)
		if err != nil {
			return "", err
		}

		addresses := make([]common.Address, 0)
		addresses = append(addresses, common.HexToAddress("0x55d398326f99059fF775485246999027B3197955"), common.HexToAddress("0xfAd476cd33Ed9213ED0a2F4c20f6865A98bf0a8B"))
		amount, _ := new(big.Int).SetString(strAmount, 10)

		bals, err := instance.GetAmountsOut(&bind.CallOpts{}, amount, addresses)
		if err != nil {
			fmt.Println(err)
			if 0 == i%4 {
				url1 = "https://bsc-dataseed4.binance.org"
			} else if 1 == i%4 {
				url1 = "https://bsc-dataseed1.binance.org"
			} else if 2 == i%4 {
				url1 = "https://bsc-dataseed.binance.org"
			} else if 3 == i%4 {
				url1 = "https://bsc-dataseed3.binance.org"
			}
			continue
		}
		balString = bals[1].String()
		break
	}

	return balString, nil
}

func GetAmountOut1(strAmount string) (string, error) {

	var balString string
	url1 := "https://bnb-bscnews.rpc.blxrbdn.com/"

	client, err := ethclient.Dial(url1)
	if err != nil {
		return "", err
	}

	tokenAddress := common.HexToAddress("0x10ED43C718714eb63d5aA57B78B54704E256024E")
	instance, err := NewPancakerouterv2(tokenAddress, client)
	if err != nil {
		return "", err
	}

	addresses := make([]common.Address, 0)
	addresses = append(addresses, common.HexToAddress("0x55d398326f99059fF775485246999027B3197955"), common.HexToAddress("0x538ac017aa01ba9665052660ea5783ba91a48092"))
	amount, _ := new(big.Int).SetString(strAmount, 10)

	bals, err := instance.GetAmountsOut(&bind.CallOpts{}, amount, addresses)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	balString = bals[1].String()
	return balString, nil
}

type eth struct {
	CoinId string
	Usd    float64
}

func requestHbsResult() (float64, error) {
	//apiUrl := "https://api-testnet.bscscan.com/api"
	apiUrl := "https://be.api.hbsswap.com/market/coin/rates"
	// URL param
	data := url.Values{}

	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		return 0, err
	}
	u.RawQuery = data.Encode() // URL encode
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(u.String())
	if err != nil {
		return 0, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var i struct {
		Data []*eth `json:"Data"`
	}
	err = json.Unmarshal(b, &i)
	if err != nil {
		return 0, err
	}

	var price float64
	for _, v := range i.Data {
		if "HBS(BEP20)" == v.CoinId { // 接收者
			price = v.Usd
		}
	}

	return price, err
}

package handlers
import(
	"encoding/json"
	"net/http"
	"junkgocode/api"
	"junkgocode/internal/tools"
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/schema"
)

func GetBankBalance(w http.ResponseWriter, r *http.Request){
var params = api.User{}
var decoder *schema.Decoder = schema.NewDecoder()
var err error
err = decoder.Decode(&params, r.URL.Query())

if err != nil{
	log.Error(err)
	api.InternalErrorHandler(w)
	return
}

var database *tools.DatabaseInterface

database, err = tools.NewDatabase()
if err != nil{
	api.InternalErrorHandler(w)
	return
}
tokenDetails := (*database).GetBankBalance(params.Username)
if tokenDetails == nil {
	log.Error(err)
	api.InternalErrorHandler(w)
	return
}

var response = api.BankBalance{
	BankBalance: (*tokenDetails).Balance,
	Code: http.StatusOK,
}
w.Header().Set("Content-Type", "application/json")
err = json.NewEncoder(w).Encode(response)
if err != nil{
	log.Error(err)
	api.InternalErrorHandler(w)
	return
}

}
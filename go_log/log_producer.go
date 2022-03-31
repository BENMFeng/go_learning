package main
import (
    "github.com/golang/glog"
    "time"
)
func main(){
    timer := time.NewTicker(1 * time.Second)
    for {
        select {
        case timeInfo:=<-timer.C:
            glog.Infof("打个日志,%s",timeInfo.Format("2006-01-02 15:04:05"))
        }
    }
}

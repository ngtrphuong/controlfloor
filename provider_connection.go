package main

import (
    //"fmt"
    //ecies "github.com/ecies/go"
    uj "github.com/nanoscopic/ujsonin/v2/mod"
)

type ProviderConnection struct {
    provChan chan ProvBase
    reqTracker *ReqTracker
}

func NewProviderConnection( provChan chan ProvBase ) (*ProviderConnection) {
    self := &ProviderConnection{
        provChan: provChan,
        reqTracker: NewReqTracker(),
    }
    
    return self
}

func (self *ProviderConnection) doPing() {
    ping := &ProvPing{
        onRes: func( root uj.JNode, raw []byte ) {
            //text := root.Get("text").String()
            //fmt.Printf("pong text %s\n", text )
        },
    }
    self.provChan <- ping
}

func (self *ProviderConnection) doClick( udid string, x int, y int, onDone func( uj.JNode, []byte ) ) {
    click := &ProvClick{
        udid: udid,
        x: x,
        y: y,
        onRes: onDone,
    }
    self.provChan <- click
}

func (self *ProviderConnection) doHardPress( udid string, x int, y int ) {
    click := &ProvHardPress{
        udid: udid,
        x: x,
        y: y,
    }
    self.provChan <- click
}

func (self *ProviderConnection) doLongPress( udid string, x int, y int ) {
    click := &ProvLongPress{
        udid: udid,
        x: x,
        y: y,
    }
    self.provChan <- click
}

func (self *ProviderConnection) doHome( udid string, onDone func( uj.JNode, []byte ) ) {
    home := &ProvHome{
        udid: udid,
        onRes: onDone,
    }
    self.provChan <- home
}

func (self *ProviderConnection) doSource( udid string, onDone func( uj.JNode, []byte ) ) {
    source := &ProvSource{
        udid: udid,
        onRes: onDone,
    }
    self.provChan <- source
}

func (self *ProviderConnection) doWifiIp( udid string, onDone func( uj.JNode, []byte ) ) {
    action := &ProvWifiIp{
        udid: udid,
        onRes: onDone,
    }
    self.provChan <- action
}

func (self *ProviderConnection) doShutdown( onDone func( uj.JNode, []byte ) ) {
    msg := &ProvShutdown{
        onRes: onDone,
    }
    self.provChan <- msg
}

func (self *ProviderConnection) doKeys( udid string, keys string, curid int, prevkeys string, onDone func( uj.JNode, []byte ) ) {
    action := &ProvKeys{
        udid: udid,
        keys: keys,
        curid: curid,
        prevkeys: prevkeys,
        onRes: onDone,
    }
    self.provChan <- action
}

func (self *ProviderConnection) doSwipe( udid string, x1 int, y1 int, x2 int, y2 int, delay float64, onDone func( uj.JNode, []byte ) ) {
    swipe := &ProvSwipe{
        udid: udid,
        x1: x1,
        y1: y1,
        x2: x2,
        y2: y2,
        delay: delay,
        onRes: onDone,
    }
    self.provChan <- swipe
}

func (self *ProviderConnection) startImgStream( udid string ) {
    self.provChan <- &ProvStartStream{ udid: udid }
}

func (self *ProviderConnection) stopImgStream( udid string ) {
    self.provChan <- &ProvStopStream{ udid: udid }
}
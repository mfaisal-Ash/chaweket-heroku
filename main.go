package main

import (
    "fmt"
    "github.com/gorilla/websocket"
    "github.com/novalagung/gubrak/v2"
    "io/ioutil"
    "log"
    "net/http"
    "strings"
)

type M map[string]interface{}

const MESSAGE_NEW_USER = "New User"
const MESSAGE_CHAT = "Chat"
const MESSAGE_LEAVE = "Leave"

var connections = make([]*WebSocketConnection, 0)
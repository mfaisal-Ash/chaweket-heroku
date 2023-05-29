package main

import (
    
    "github.com/sidiq200/chaweket-heroku/typestruct"
    "github.com/sidiq200/chaweket-heroku/module"
    

)

type M map[string]interface{}

const MESSAGE_NEW_USER = "New User"
const MESSAGE_CHAT = "Chat"
const MESSAGE_LEAVE = "Leave"

var connections = make([]*WebSocketConnection, 0)
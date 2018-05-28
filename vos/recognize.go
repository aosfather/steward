package vos
/**
语音识别
*/

//根据给定的语音，识别出文本
type RecognizeFunction func(data []byte) string

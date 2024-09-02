package auth

// this is private function => only using small First Letter => s
func sessionInfo() string {
	return "Main private funcion, This can be use only in this package"
}

// this is public function => only using caps First Letter => G
func GetSession() string {
	return sessionInfo()
}

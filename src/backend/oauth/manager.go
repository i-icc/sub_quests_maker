package oauth

type Manager struct {
	tokens map[string]*Token
}

var Mg Manager

func Init() {
	Mg.tokens = map[string]*Token{}
}

func (m *Manager) Save(token *Token) string {
	m.tokens[token.id] = token

	return token.id
}

func (m *Manager) Exists(tokenId string) bool {
	_, r := m.tokens[tokenId]
	return r
}

func (m *Manager) Destroy(tokenId string) {
	delete(m.tokens, tokenId)
}

func (m *Manager) GetToken(tokenId string) *Token {
	return m.tokens[tokenId]
}

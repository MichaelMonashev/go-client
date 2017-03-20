// Package code contain command ids for clent-server interaction.
package code

// коды запросов от клиентов
const (
	CONNECT    = iota // запрос номера ноды и номера соединения
	PING              // запрос для проверки прохождения пакетов до сервера
	TOUCH             // запрос для обновления времени последней активности клиента
	DISCONNECT        // запрос на отсоединение

	LOCK      // запрос заблокировать ключ
	UPDATE    // запрос обновить информацию о заблокированном ключе
	UNLOCK    // запрос снять блокировку с ранее заблокированного ключа
	UNLOCKALL // запрос снять все блокировки
)

// коды ответов клиентам
const (
	OPTIONS = iota // ответ на CONNECT-запрос. Содержит номер сервера, номер соединения с ней и список всех серверов, если их несколько.
	PONG           // ответ на PING-запрос

	// коды ответов на DISCONNECT, TOUCH, LOCK, UPDATE, UNLOCK и UNLOCKALL
	OK           // ответ: всё хорошо
	DISCONNECTED // ответ: клиент давно не делал запросы, его данные проэкспайрились, и поэтому ему нужно снова соединиться
	ISOLATED     // ответ: клиент был изолирован сервером. Ему нужно перестать работать с данными
	LOCKED       // ответ: ключ заблокирован кем-то ещё
	//	WORKING      // ответ: запрос обрабатывается (т.е. клиенту надо обновить таймаут на ожидание ответа);
	REDIRECT // ответ: повторить запрос на другую ноду
	ERROR    // ответ: ошибка
)

// флаги транспортного уровня
const (
	FRAGMENTED    = 1 << iota // один из нескольких фрагментов
	LAST_FRAGMENT             // последний фрагмент
	BUSY                      // сервер перегружен, присылайте меньше запросов
)
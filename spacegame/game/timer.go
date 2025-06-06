package game

type Timer struct {
	currentTicks int
	targetTicks int
}

func NewTimer(targetTicks int) *Timer {
	return &Timer{
		currentTicks: 0,
		targetTicks: targetTicks,
	}
}

func (t *Timer) Update(){
	if t.currentTicks < t.targetTicks {
		t.currentTicks++
	}
}

func (t *Timer) isReady() bool {
	return t.currentTicks >= t.targetTicks
}

func (t *Timer) resetTimer() {
	t.currentTicks = 0
}
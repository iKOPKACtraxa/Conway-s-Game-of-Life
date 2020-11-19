		a.Show()   //displays one screen of current Universe (one frame)
		Step(a, b) //generates next step
		a, b = b, a
		time.Sleep(time.Second / 16)
	}
}

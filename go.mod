module helloworld

go 1.16

require (
	rsc.io/quote v1.5.2
	tenyunops.com/nashyang/calculator v0.0.0
)

replace tenyunops.com/nashyang/calculator => ../calculator

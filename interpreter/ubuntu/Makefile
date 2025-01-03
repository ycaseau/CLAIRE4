# +-------------------------------------------------------------------------------+
# | Part 3: launch automated tests                                                |
# + ------------------------------------------------------------------------------+

# non regression tests
test:	claire4
	claire4 -m bu1
	claire4 -m bu2
	claire4 -m bu3
	claire4 -m bu4
	claire4 -m bu5
	claire4 -m bu6
	claire4 -m bu7
	claire4 -m bu8
	claire4 -m bu9
	claire4 -m bu10
	claire4 -m bu11
	claire4 -m bu12
	claire4 -m bu13
	claire4 -m bu14
	claire4 -m bu15	
	claire4 -m bu16

ctest:	claire4
	claire4 -O -cm bu1; bu1
	claire4 -cm bu2; bu2
	claire4 -O -cm bu3; bu3
	claire4 -O -cm bu4; bu4
	claire4 -O -cm bu5; bu5
	claire4 -O -cm bu6; bu6 
	claire4 -O -cm bu7; bu7
	claire4 -O -cm bu8; bu8
	claire4 -O -cm bu9; bu9
	claire4 -O -cm bu10; bu10
	claire4 -O -cm bu11; bu11
	claire4 -O -cm bu12; bu12
	claire4 -O -cm bu13; bu13
	claire4 -O -cm bu14; bu14
	claire4 -O -cm bu15; bu15	
	claire4 -O -cm bu16; bu16


# performance tests
perf:	claire4
	claire4 -mx mFib 
	claire4 -mx mList 
	claire4 -mx mSet 
	claire4 -mx mDict 
	claire4 -mx mObj 
	claire4 -mx mSend 

cperf:	claire4
	claire4 -cx mFib
	claire4 -cx mList
	claire4 -cx mSet
	claire4 -cx mDict
	claire4 -cx mObj 
	claire4 -cx mSend
	mFib
	mList
	mSet
	mDict
	mObj
	mSend

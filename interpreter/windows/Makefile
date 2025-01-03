###################################################################################
# This is the global makefile for CLAIRE1 : first version of CLAIRE4 interpreter  #
###################################################################################

# warning : the place of the comment is critical if we do not want to add a whitespace to the var


# root for claire 4
ROOT = E:\claire\v4.1

# where we are - last step
WHERE = $(ROOT)\go

# go files
SRC = $(WHERE)\src

#  claire files are on the root GitHub directory
DROP = $(ROOT)



# +-------------------------------------------------------------------------------+
# | Part 2: make claire1 interpreter + claire4 compiler                           |
# + ------------------------------------------------------------------------------+

# sources for Kernel files
DGK = $(ROOT)\go\src\Kernel
K4FILES = $(DGK)\ClKernel.go $(DGK)\ClReflect.go $(DGK)\ClBag.go $(DGK)\ClString.go $(DGK)\ClEnv.go $(DGK)\ClUtil.go

#files to build claire4 from claire4
CSRC = $(DROP)\meta
CSRC2 = $(DROP)\compile
CFILES = $(CSRC)\method.cl $(CSRC)\object.cl $(CSRC)\function.cl $(CSRC)\types.cl
LFILES = $(CSRC)\pretty.cl $(CSRC)\call.cl $(CSRC)\control.cl $(CSRC)\define.cl
RFILES = $(CSRC)\read.cl $(CSRC)\syntax.cl $(CSRC)\file.cl $(CSRC)\inspect.cl
OFILES = $(CSRC2)\osystem.cl $(CSRC2)\otool.cl $(CSRC2)\ocall.cl $(CSRC2)\ocontrol.cl $(CSRC2)\odefine.cl
GFILES = $(CSRC2)\gosystem.cl $(CSRC2)\gogen.cl $(CSRC2)\goexp.cl $(CSRC2)\gostat.cl $(CSRC2)\gomain.cl

# module based recompilation
CORE = $(ROOT)\go\src\Core\Core-meta.go
LANGUAGE = $(ROOT)\go\src\Language\Language-meta.go
READER = $(ROOT)\go\src\Reader\Reader-meta.go
OPTIMIZE = $(ROOT)\go\src\Optimize\Optimize-meta.go
GENERATE = $(ROOT)\go\src\Generate\Generate-meta.go

$(CORE):	$(CFILES) sclaire4.exe
	sclaire4 -cc Core

$(LANGUAGE):	$(LFILES) sclaire4.exe
	sclaire4 -cc Language

$(READER):	$(RFILES) sclaire4.exe
	sclaire4 -cc Reader

$(OPTIMIZE):	$(CFILES) sclaire4.exe
	sclaire4 -cc Optimize

$(GENERATE):	$(GFILES) sclaire4.exe
	sclaire4 -cc Generate

test: $(LANGUAGE)	

#create a claire1 for debug
claire1:	$(DGK) $(CORE) $(LANGUAGE) $(READER) 
	sclaire4 -sf Reader -o claire1
	go build src\claire1.go

#compilation scripts : claire1 is a debug light version, claire4 is the real thing
claire4: claire1 $(OPTIMIZE) $(GENERATE) init.cl
	sclaire4 -sf Generate -o claire4
	go build src\claire4.go
	@echo "claire4 executable built"

# clean all generated go files
clean1:
	rm src\Core\*.go
	rm src\Language\*.go
	rm src\Reader\*.go


step:	claire4	
	cp claire4 sclaire4
	

# +-------------------------------------------------------------------------------+
# | Part 3: launch automated tests                                                |
# + ------------------------------------------------------------------------------+

# non regression tests
test:	claire4.exe
	rm -f log3
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

ctest:	claire4.exe
	rm -f clog3
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

# default safety in init = maximal compilation
ctesto:	claire4.exe
	rm -f clog3
	claire4 -cm bu1; bu1
	claire4 -cm bu2; bu2
	claire4 -O -cm bu3; bu3
	claire4 -O2 -cm bu4; bu4
	claire4 -cm bu5; bu5
	claire4 -cm bu6; bu6 
	claire4 -cm bu7; bu7
	claire4 -cm bu8; bu8
	claire4 -cm bu9; bu9
	claire4 -cm bu10; bu10
	claire4 -cm bu11; bu11
	claire4 -cm bu12; bu12
	claire4 -cm bu13; bu13
	claire4 -cm bu14; bu14
	claire4 -cm bu15; bu15	
	claire4 -cm bu16; bu16


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

	

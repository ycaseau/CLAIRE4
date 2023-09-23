// --- System configuration file for "claire4" , ["Saturday 09-23-2023"] ---

package main
import (
	. "Kernel"
	"Core"
	"Language"
	"Reader"
	"Optimize"
	"Generate"
)

//load function : create and load modules
func Load() { 
  It := C_claire
  //module definitions 
  Core.It = InitModule("Core",C_mClaire,MakeConstantList(C_Kernel.Id()),
  	"/Users/ycaseau/Dropbox/src/clairev4.10/src/meta",
	MakeList(ToType(C_string.Id()),MakeString("method").Id(),
    MakeString("object").Id(),
    MakeString("function").Id(),
    MakeString("types").Id()))
  Language.C_iClaire = InitModule("iClaire",It,ToType(C_module.Id()).EmptyList(),
  	"",
	ToType(C_string.Id()).EmptyList())
  Language.It = InitModule("Language",Language.C_iClaire,MakeConstantList(C_Kernel.Id(),Core.It.Id()),
  	"/Users/ycaseau/Dropbox/src/clairev4.10/src/meta",
	MakeList(ToType(C_string.Id()),MakeString("pretty").Id(),
    MakeString("call").Id(),
    MakeString("control").Id(),
    MakeString("define").Id()))
  Reader.It = InitModule("Reader",Language.C_iClaire,MakeConstantList(C_Kernel.Id(),Core.It.Id(),Language.It.Id()),
  	"/Users/ycaseau/Dropbox/src/clairev4.10/src/meta",
	MakeList(ToType(C_string.Id()),MakeString("read").Id(),
    MakeString("syntax").Id(),
    MakeString("file").Id(),
    MakeString("inspect").Id()))
  Optimize.C_Compile = InitModule("Compile",Language.C_iClaire,MakeConstantList(C_mClaire.Id()),
  	"",
	ToType(C_string.Id()).EmptyList())
  Optimize.It = InitModule("Optimize",Optimize.C_Compile,MakeConstantList(C_mClaire.Id(),Reader.It.Id()),
  	"/Users/ycaseau/Dropbox/src/clairev4.10/src/compile",
	MakeConstantList(MakeString("osystem").Id(),
    MakeString("otool").Id(),
    MakeString("ocall").Id(),
    MakeString("ocontrol").Id(),
    MakeString("odefine").Id()))
  Generate.It = InitModule("Generate",Optimize.C_Compile,MakeConstantList(C_mClaire.Id(),Optimize.It.Id()),
  	"/Users/ycaseau/Dropbox/src/clairev4.10/src/compile",
	MakeConstantList(MakeString("gosystem").Id(),
    MakeString("gogen").Id(),
    MakeString("goexp").Id(),
    MakeString("gostat").Id(),
    MakeString("gomain").Id()))
  
  // module load 
  Core.MetaLoad()
  Language.MetaLoad()
  Reader.MetaLoad()
  Optimize.MetaLoad()
  Generate.MetaLoad()
  ClEnv.Module_I = Generate.It; 
  } 

// reboot function is created by compiler 
func E_reboot(s EID) EID { 
  Bootstrap()
  Load()
  Reader.C_reader.Fromp = ClEnv.Cin;
  ClEnv.Module_I = C_claire
  return EVOID
  } 


// the main function 
func main() { 
  MemoryFlags()
  E_reboot(EVOID)
  ToMethod(C_reboot.Restrictions.ValuesO()[0]).Functional = MakeFunction1(E_reboot,"reboot")
  Reader.C_reader.Fromp = ClEnv.Cin
  Generate.F_Generate_complex_main_void()
  } 

To install under SBCL/Windows *without* TCL/tk
____________________________________________

Start SBCL

Enter (load "install.lsp").  

The executable sbcl-shen.exe will be created.

To install under SBCL/Windows *with* TCL/tk
____________________________________________

To run Shen/tk under Windows; install TCL/tk.  

Enter (load "install-tk.lsp").

The executable sbcl-shen+tk.exe will be created.

Edit the file shen-tk.bat and replace the pathnames in 
the second line to those suitable for your installation.
**Click on this batch file to invoke Shen/tk.**

Now go to Lib and look in the file Tk/root.tcl.  You will see two lines.

set in {C:/Users/shend/OneDrive/Desktop/Shen/S39/shen-to-tcl.txt} 
set out {C:/Users/shend/OneDrive/Desktop/Shen/S39/tcl-to-shen.txt}

Change the pathnames to your installation.

IMPORTANT!
*******************************************************
Note before quitting Shen/tk you should click on the 
exit button in the IDE and confirm you want to exit 
before shutting down the command window and root window.  
This will give a clean shutdown.
*******************************************************

Failure to do this will leave a zombie TCL process that 
will definitely interfere with your next session.  In 
that event you will have to kill that process in the task 
manager.
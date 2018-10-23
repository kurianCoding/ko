# knock out testing 

this is a  command line application that lets you add and remove
tests. All you have to do is write your test functions the way
you would in a normal file and then rename the file to test.ko

now if you run
    `cmd <pathtopackage>` 
it will create a test file that contains all the test functions
in the test.ko.

So, why use this instead of `go test`?  with this you can choose
which of those test functions are tested by including a `.testignore`
function in your package directory. 

If you have a testfunction called `func TestThis(t *testing.T){}` in 
test.ko which you don't want to include right now, simply add a line
`TestThis` in `.testignore`.


### 1st aasembely input

> using System.Diagnostics;

### now declear object on button (` public partial class Form1 : Form `)

```
private Process process=new Process();
private processStartInfo info =new ProcessStartInfo();
```
###  now initialize on (  public Form1(){ )
> InitializeComponent();
```         info.WindowStyle = ProcessWindowStyle.Hidden;
            info.CreateNoWindow = true;
            info.UseShellExecute= false;
            info.RedirectStandardOutput = true;
```

### Now goto event button & added folder on `bin/debug/ToolsName`
```
 try
            {
                info.FileName = "Name//adb.exe";
                info.Arguments = "devices";
                process.StartInfo = info;

                process.Start();

                textBox1.Text = process.StandardOutput.ReadToEnd();

            }
            catch(Exception ex)
            {
                MessageBox.Show(ex.Message);
            }
```

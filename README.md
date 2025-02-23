# 📄 **Attendance PDF Generator** 🖨️

## 📌 **Overview** 🚀
A **simple and efficient** Go program that **automatically generates a PDF report** 📑 containing **student attendance records**! With the help of the **gopdf** library, it structures the data into a **neat and readable table**, featuring **serial numbers, USN, names, login times**, and **logout times**.

---

## ✨ **Features** 🎉
- ✅ **Automatically generates a PDF** with detailed attendance data 📜
- ✅ **Uses gopdf** for seamless PDF rendering 🖋️
- ✅ **Neatly formatted table** with proper alignment 🏫
- ✅ **Handles multiple students & long names** without issues 📏

---

## 🛠️ **Tech Stack & Dependencies** 💻
**Built with:**
- 🔹 **Go** – High-performance backend programming language 🚀
- 🔹 **gopdf** – A powerful and flexible library for creating PDFs 📄

---

## 📌 **Before Running the Program** 💡
Ensure you have the following installed:

- ✅ **Go** (Check version: `go version`)
- ✅ **gopdf** package:

```sh
go get github.com/signintech/gopdf



## ⚙️ Configuration  
-------------------->

📌 Modify the **Students array** in the code to include your own **USN, Name, Login, and Logout details**.

---

## 📜 Example Output 📑  
---------------------->

🎯 The generated PDF will contain a **structured attendance table** in the following format:

|  Sr. No. | 🎓 USN     |     🏷️ Name       |   ⏰ Login    | ⏳ Logout   |
|    --    |    ---     |         ---       |      ---      |       ---   |
|    1️⃣    | 4AL21EC053 | **John Doe**      |   08:30 AM    |   12:30 PM  |
|    2️⃣    | 4AL21EC054 | **Michael Smith** |   09:00 AM    |   01:00 PM  |

---

## ⚠️ Notes 📝  
-------------->

⚡ **Long names may get truncated** in the PDF, adjust formatting if needed  
⚡ Ensure **gopdf** is installed before running the script

------>


🚀 **Happy coding!** 🚀


import sys
import display
import urllib2
import urllib

from PyQt4 import QtGui, uic

qtCreatorFile = "Login.ui"

Ui_MainWindow, QtBaseClass = uic.loadUiType(qtCreatorFile)


class Login(QtGui.QMainWindow, Ui_MainWindow):
    IP = ""
    
    def __init__(self):
        QtGui.QMainWindow.__init__(self)
        Ui_MainWindow.__init__(self)
        self.setupUi(self)
        import socket

        hostname = socket.gethostname()
        self.IP = socket.gethostbyname(hostname)
        print(self.IP)

        self.setWindowTitle("Login")
        self.line.setStyleSheet('background-color:#111111')
        self.lineEdit.setPlaceholderText("     Student ID")
        self.lineEdit_2.setPlaceholderText("     Password")
        self.lineEdit_2.setEchoMode(QtGui.QLineEdit.Password)
        self.label_3.setStyleSheet('color: red')
        self.label_2.setStyleSheet('color: black')
        self.label_3.setStyleSheet('color: red')
        self.label_5.setStyleSheet('background-image:url("id.png");')
        self.label_6.setStyleSheet('background-image:url("lock.png");')
        size = self.geometry()
        self.frame.setStyleSheet('border-radius:4px;')
        self.label_12.setStyleSheet('background-image:url("fb.png");')
        self.label_13.setStyleSheet('background-image:url("gplus.png");')
        self.label_14.setStyleSheet('background-image:url("tweet.png");')
        screen = QtGui.QDesktopWidget().screenGeometry(self)
        self.move((screen.width() - size.width()) / 2, (screen.height() - size.height()) / 2)
        self.lineEdit.setStyleSheet('border-radius:4px;color: black;border:1px solid black;')
        self.lineEdit_2.setStyleSheet('border-radius:4px;color: black;border:1px solid black;')
        self.label.setStyleSheet('color:#111111')
        self.pushButton.setStyleSheet('border-radius:3px;box-shadow: 0 2px 0 0 #1e7b1e;background-color: #1e7b1e;width: 100%;padding: 10px;color: #ffffff;')

        self.pushButton.clicked.connect(self.pushButton_clicked)

    def pushButton_clicked(self):
        ID = self.lineEdit.text()
        Password = self.lineEdit_2.text()
        self.label_3.setText("")
        print(str(Password))
        if str(ID).strip() == "" or str(Password).strip() == "" or (ID == "" and Password == ""):
            self.label_3.setText("Please fill both the fields")
        else:
                parameters = {"username": ID, "password": Password}
                data = urllib.urlencode(parameters)
                try:
                    req = urllib2.Request("http://localhost:8080/stuLogin/", data)
                    response = urllib2.urlopen(req, timeout=4)
                    name = response.read()
                    print(name)
                    if name != "INVALID":
                        self.label_3.setText("")
                        self.call = display.Display(ID, name, self.IP)
                        self.call.show()
                        self.close()

                    else:
                        self.label_3.setText("Username or Password is invalid")
                except urllib2.URLError:
                    self.label_3.setText("Connection error. Try again")

                 
if __name__ == '__main__':

    app = QtGui.QApplication(sys.argv)
    window = Login()
    window.show()
    sys.exit(app.exec_())

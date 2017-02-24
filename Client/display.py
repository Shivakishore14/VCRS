import urllib
import login
import urllib2
import json

from PyQt4 import QtCore, QtGui, uic

qtCreatorFile = "questions.ui"

Ui_MainWindow, QtBaseClass = uic.loadUiType(qtCreatorFile)


class Display(QtGui.QMainWindow, Ui_MainWindow):
    ID = ""
    IP = ""
    id = 1
    Qtype = ""
    total = 0
    j = {}
    answers = {}
    test = {}
    selectedTest = ""

    def __init__(self, ID, name, IP):
        QtGui.QMainWindow.__init__(self)
        Ui_MainWindow.__init__(self)
        self.setupUi(self)
        self.setWindowTitle("Test Session")
        self.ID = str(ID)
        self.IP = IP

        self.label.setText("Hi " + name)
        self.toolButton.setVisible(False)
        self.toolButton_2.setVisible(False)
        self.toolButton.setStyleSheet('background-color:DeepSkyBlue')
     

        self.label.setStyleSheet('color:#111111')
        self.label_2.setStyleSheet('color:#111111')
        self.label_4.setStyleSheet('color:#111111')
        self.label_8.setStyleSheet('color:#111111')
        self.label_5.setStyleSheet('color:#111111')
        self.label_7.setStyleSheet('color:#111111')
        self.line.setStyleSheet('background-color:#111111')
        self.line_2.setStyleSheet('background-color:#111111')
        self.line_2.setVisible(False)
        self.label_6.setStyleSheet('color:#111111')
        self.label3.setStyleSheet('color:#111111')
        self.pushButton.setStyleSheet(
            'background-color: #1e7b1e;width:50px;height:50px;line-height:50px;color:#ffffff')
        self.pushButton_2.setStyleSheet(
            'background-color: #1e7b1e;width:50px;height:50px;line-height:50px;color:#ffffff')
        self.pushButton_3.setVisible(False)
        self.pushButton_3.setStyleSheet(
            'background-color: #1e7b1e;width:50px;height:50px;line-height:50px;color:#ffffff')
        self.toolButton_2.setStyleSheet('background-color:DeepSkyBlue')
        self.radioButton1.setVisible(False)
        self.radioButton2.setVisible(False)
        self.radioButton3.setVisible(False)
        self.radioButton4.setVisible(False)
        self.pushButton_2.setVisible(False)
        self.scrollArea_2.setVisible(False)
        self.scrollArea_4.setVisible(False)
        self.scrollArea.setVisible(False)

        self.label_4.setVisible(False)
        self.label_5.setVisible(False)
        self.label_6.setVisible(False)

        self.label3.setVisible(False)
        self.label_7.setVisible(False)
        self.pushButton.clicked.connect(self.showTests)
        self.tableWidget.setVisible(False)
        header = ["Test Name"]
        self.tableWidget.setHorizontalHeaderLabels(header)
        self.tableWidget.cellClicked.connect(self.cell_was_clicked)
        self.pushButton_3.clicked.connect(self.check)
        self.show()

    def showTests(self):
        self.call = MyPopup1()
        self.call.doit1()
        parameters = {"getTest": self.ID}
        data = urllib.urlencode(parameters)
        try:
            req = urllib2.Request("http://localhost:8080/getData/", data)
            response = urllib2.urlopen(req, timeout=6)
        except urllib2.URLError, e:
            raise Exception("There was an error: %r" % e)
        jsonvalue = response.read()
        self.test = json.loads(jsonvalue)
        self.total = 0
        self.pushButton.setVisible(False)
        self.pushButton_2.setVisible(True)
        self.tableWidget.setVisible(True)
        self.label_2.setText("The tests you have today : ")
        for i in range(len(self.test)):
            self.tableWidget.insertRow(i)
            self.tableWidget.setItem(i, 0, QtGui.QTableWidgetItem(self.test[i]["name"].strip()))

        self.pushButton_2.setText("Ready")

    def cell_was_clicked(self, row, column):
        print(row,column)
        item = self.test[row]["name"].strip()
        print(item)
        self.selectedTest = self.test[row]["name"]
        self.total = self.test[row]["no"]
        self.pushButton_2.clicked.connect(self.start)

    def start(self):
        parameters = {"setTest": self.selectedTest}
        data = urllib.urlencode(parameters)
        try:
            req = urllib2.Request("http://localhost:8080/getData/", data)
            response = urllib2.urlopen(req, timeout=4)
        except urllib2.URLError, e:
            raise Exception("There was an error: %r" % e)
        jsonvalue = response.read()
        print(jsonvalue)
        if str(jsonvalue) == "OK":
            self.pushButton_2.setText("Start")
            self.pushButton_2.clicked.disconnect(self.start)

            self.pushButton_2.clicked.connect(self.onClicked)

    def onClicked(self):
        parameters = {"id": self.id, "testName": self.selectedTest}
        print(self.selectedTest)
        data = urllib.urlencode(parameters)
        try:
            req = urllib2.Request("http://localhost:8080/getQuestions/", data)
            response = urllib2.urlopen(req, timeout=4)
        except urllib2.URLError, e:
            raise Exception("There was an error: %r" % e)
        jsonvalue = response.read()
        print(str(jsonvalue).strip() )
        if str(jsonvalue).strip()[0] is "N":
            self.label_8.setText("The test you selected is not available")
        else:
            self.tableWidget.setVisible(False)
            self.pushButton.setVisible(False)
            self.label_8.setVisible(False)
            self.label3.setVisible(True)
            self.toolButton.setVisible(True)
            self.toolButton_2.setVisible(True)
            self.radioButton1.setVisible(True)
            self.radioButton2.setVisible(True)
            self.radioButton3.setVisible(True)
            self.radioButton4.setVisible(True)
            self.scrollArea_2.setVisible(True)
            self.scrollArea_4.setVisible(True)
            self.scrollArea.setVisible(True)
            self.label_4.setVisible(True)
            self.label_5.setVisible(True)
            self.label_6.setVisible(True)
            self.label_7.setVisible(True)
            self.line_2.setVisible(True)

            self.j = json.loads(jsonvalue)
            self.label_2.setText("Your questions for today :  " + str(self.id) + "/" + str(self.total))
            self.label3.setText( self.j["id"] + ".  " + self.j["question"])
            self.Qtype = self.j["type"]
            if self.Qtype == "text":
                self.radioButton1.setText("A : ")
                self.label_4.setText(self.j["option1"])
                self.radioButton2.setText("B : ")
                self.label_6.setText(self.j["option2"])
                self.radioButton3.setText("C : ")
                self.label_5.setText(self.j["option3"])
                self.radioButton4.setText("D : ")
                self.label_7.setText(self.j["option4"])
            else:
                self.label_4.setOpenExternalLinks(True)
                self.label_5.setOpenExternalLinks(True)
                self.label_6.setOpenExternalLinks(True)
                self.label_7.setOpenExternalLinks(True)
                self.radioButton1.setText("A : ")

                self.label_4.setText("Option 1")
                self.label_4.mousePressEvent = self.repeat
                self.radioButton2.setText("B : ")
                self.label_6.setText("Option 2")
                self.label_6.mousePressEvent = self.repeat2
                self.radioButton3.setText("C : ")

                self.label_5.setText("Option 3")
                self.label_5.mousePressEvent = self.repeat3
                self.radioButton4.setText("D : ")

                self.label_7.setText("Option 4")
                self.label_7.mousePressEvent = self.repeat4

            self.toolButton_2.clicked.connect(self.move)
            self.toolButton.clicked.connect(self.back)
            self.pushButton_3.setVisible(True)

    def move(self):
        if self.id  <  int(self.total.strip()):
            if self.radioButton1.isChecked():
                self.answers[self.id-1] = "option1"
                self.radioButton1.setChecked(False)

            if self.radioButton2.isChecked():
                self.answers[self.id - 1] = "option2"
                self.radioButton2.setChecked(False)

            if self.radioButton3.isChecked():
                self.answers[self.id - 1] = "option3"
                self.radioButton3.setChecked(False)

            if self.radioButton4.isChecked():
                self.answers[self.id - 1] = "option4"
                self.radioButton4.setChecked(False)
            self.id += 1

            parameters = {"id": self.id, "testName": self.selectedTest}
            data = urllib.urlencode(parameters)
            self.label_2.setText("Your questions for today :   " + str(self.id) + "/" + str(self.total))
            try:
                req = urllib2.Request("http://localhost:8080/getQuestions/", data)
                response = urllib2.urlopen(req, timeout=4)
            except urllib2.URLError, e:
                raise Exception("There was an error: %r" % e)
            jsonvalue = response.read()
            self.j = json.loads(jsonvalue)
            self.label3.setText(self.j["id"]+".  "+self.j["question"])
            self.Qtype = self.j["type"]

            if self.Qtype == "text":
                self.radioButton1.setText("A : ")
                self.label_4.setText(self.j["option1"])
                self.radioButton2.setText("B : ")
                self.label_6.setText(self.j["option2"])
                self.radioButton3.setText("C : ")
                self.label_5.setText(self.j["option3"])
                self.radioButton4.setText("D : ")
                self.label_7.setText(self.j["option4"])
            else:
                self.label_4.setOpenExternalLinks(True)
                self.label_5.setOpenExternalLinks(True)
                self.label_6.setOpenExternalLinks(True)
                self.label_7.setOpenExternalLinks(True)
                self.radioButton1.setText("A : ")

                self.label_4.setText("Option 1")
                self.label_4.mousePressEvent = self.repeat
                self.radioButton2.setText("B : ")
                self.label_6.setText("Option 2")
                self.label_6.mousePressEvent = self.repeat2
                self.radioButton3.setText("C : ")

                self.label_5.setText("Option 3")
                self.label_5.mousePressEvent = self.repeat3
                self.radioButton4.setText("D : ")

                self.label_7.setText("Option 4")
                self.label_7.mousePressEvent = self.repeat4

    def back(self):
        if self.id >= 2:

            if self.radioButton1.isChecked():
                self.answers[self.id - 1] = "option1"
                self.radioButton1.setChecked(False)

            if self.radioButton2.isChecked():
                self.answers[self.id - 1] = "option2"
                self.radioButton2.setChecked(False)

            if self.radioButton3.isChecked():
                self.answers[self.id - 1] = "option3"
                self.radioButton3.setChecked(False)

            if self.radioButton4.isChecked():
                self.answers[self.id - 1] = "option4"
                self.radioButton4.setChecked(False)

            self.id = self.id - 1
            parameters = {"id": self.id, "testName": self.selectedTest}
            data = urllib.urlencode(parameters)
            self.label_2.setText("Your questions for today :   " + str(self.id) + "/" + str(self.total))
            try:
                req = urllib2.Request("http://localhost:8080/getQuestions/", data)
                response = urllib2.urlopen(req, timeout=4)
            except urllib2.URLError, e:
                raise Exception("There was an error: %r" % e)

            jsonvalue = response.read()
            self.j = json.loads(jsonvalue)
            self.label3.setText(self.j["id"]+".  "+self.j["question"])
            self.Qtype = self.j["type"]
            if self.Qtype == "text":
                self.radioButton1.setText("A : ")
                self.label_4.setText(self.j["option1"])
                self.radioButton2.setText("B : ")
                self.label_6.setText(self.j["option2"])
                self.radioButton3.setText("C : ")
                self.label_5.setText(self.j["option3"])
                self.radioButton4.setText("D : ")
                self.label_7.setText(self.j["option4"])
            else:
                self.label_4.setOpenExternalLinks(True)
                self.label_5.setOpenExternalLinks(True)
                self.label_6.setOpenExternalLinks(True)
                self.label_7.setOpenExternalLinks(True)
                self.radioButton1.setText("A : ")

                self.label_4.setText("Option 1")
                self.label_4.mousePressEvent = self.repeat
                self.radioButton2.setText("B : ")
                self.label_6.setText("Option 2")
                self.label_6.mousePressEvent = self.repeat2
                self.radioButton3.setText("C : ")

                self.label_5.setText("Option 3")
                self.label_5.mousePressEvent = self.repeat3
                self.radioButton4.setText("D : ")

                self.label_7.setText("Option 4")
                self.label_7.mousePressEvent = self.repeat4
            self.pushButton_2.setText("Submit")


    def check(self):
        confirm = QtGui.QMessageBox.question(self, 'Confirm Submission', "Are you Sure you want to Submit?", QtGui.QMessageBox.Ok, QtGui.QMessageBox.Cancel)
        if confirm == QtGui.QMessageBox.Ok:
            self.submit()

    def submit(self):
        if self.radioButton1.isChecked():
            self.answers[self.id - 1] = "option1"
            self.radioButton1.setChecked(False)

        if self.radioButton2.isChecked():
            self.answers[self.id - 1] = "option2"
            self.radioButton2.setChecked(False)

        if self.radioButton3.isChecked():
            self.answers[self.id - 1] = "option3"
            self.radioButton3.setChecked(False)

        if self.radioButton4.isChecked():
            self.answers[self.id - 1] = "option4"
            self.radioButton4.setChecked(False)

        answers = []
        for i in range(len(self.answers)):
            answers.append(self.answers[i])
        print self.selectedTest

        parameters = {"testname": self.selectedTest, "sid": self.ID, "ans": answers}

        convert = json.dumps(parameters)
        parameters2 = {"response": convert}

        data = urllib.urlencode(parameters2)
        try:
            req = urllib2.Request("http://localhost:8080/saveResponse/", data)
            response = urllib2.urlopen(req, timeout=4)
        except urllib2.URLError, e:
            raise Exception("There was an error: %r" % e)
        jsonvalue = response.read()

        if str(jsonvalue) == "OK":
            self.radioButton1.setVisible(False)
            self.radioButton2.setVisible(False)
            self.radioButton3.setVisible(False)
            self.radioButton4.setVisible(False)
            self.pushButton_2.setVisible(True)
            self.label3.setVisible(True)
            self.label_2.setVisible(False)
            self.label_4.setVisible(False)
            self.label_5.setVisible(False)
            self.label_6.setVisible(False)
            self.label_7.setVisible(False)
            self.toolButton.setVisible(False)
            self.toolButton_2.setVisible(False)
            self.pushButton_3.setVisible(False)
            self.pushButton_2.setText("Log out")
            self.label3.setText("\n \n"+"Thank You. You have completed your test successfully.")
            self.label_8.setText("Take next test")
            self.pushButton_2.clicked.disconnect(self.onClicked)
            self.pushButton_2.clicked.connect(self.logout)

    def logout(self):
        self.pushButton_2.disconnect()
        self.pushButton_3.disconnect(self.check)
        self.call = login.Login()
        self.call.show()
        self.close()

    def repeat(self, event):
        self.call = MyPopup(self.j["option1"])

        self.call.doit()

    def repeat2(self, event):
        self.call = MyPopup(self.j["option2"])

        self.call.doit()

    def repeat3(self, event):
        self.call = MyPopup(self.j["option3"])

        self.call.doit()

    def repeat4(self, event):
        self.call = MyPopup(self.j["option4"])

        self.call.doit()


class MyPopup(QtGui.QWidget):
        data = None

        def __init__(self, url):
            QtGui.QWidget.__init__(self)
            self.data = urllib2.urlopen("http://localhost:8080"+url).read()

        def disp(self):
            pixmap = QtGui.QPixmap()
            pixmap.loadFromData(self.data)
            label = QtGui.QLabel(self)
            label.setMaximumSize(600, 300)
            label.setPixmap(pixmap)

        def center(self):
            frameGm = self.frameGeometry()
            screen = QtGui.QApplication.desktop().screenNumber(QtGui.QApplication.desktop().cursor().pos())
            centerPoint = QtGui.QApplication.desktop().screenGeometry(screen).center()
            frameGm.moveCenter(centerPoint)
            self.move(frameGm.topLeft())

        def doit(self):
            print "Opening a new popup window..."
            self.disp()
            self.setGeometry(QtCore.QRect(100, 100, 600, 300))
            self.center()
            self.show()


class MyPopup1(QtGui.QWidget):
        def __init__(self):
            QtGui.QWidget.__init__(self)

        def center(self):
            frameGm = self.frameGeometry()
            screen = QtGui.QApplication.desktop().screenNumber(QtGui.QApplication.desktop().cursor().pos())
            centerPoint = QtGui.QApplication.desktop().screenGeometry(screen).center()
            frameGm.moveCenter(centerPoint)
            self.move(frameGm.topLeft())

        def disp1(self):
                label = QtGui.QLabel(self)
                label.setMaximumSize(650, 250)
                label.setStyleSheet('color:#ea4c88;font-size:14px;font-family:Lato; padding-left:5px')
                label.setText(
                    "\t\t\t\t\t\t<b><br><br>Instructions:</b><br><br>1. You can answer all the questions any number of times until you click submit button <br>or the test time expires.<br><br>\n\t2. You can change your option at any point of time.<br><br>\n\t3. Once you click the submit button you will not be allowed to change your options.<br><br>\n\n\t Do well. <b>ALL THE BEST!!!</b>")

        def doit1(self):
            self.setWindowTitle("Instructions")
            print "Opening a new popup window..."
            self.disp1()
            self.setGeometry(QtCore.QRect(100, 100, 670, 250))
            self.center()
            self.show()


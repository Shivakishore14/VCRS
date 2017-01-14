import sys
import MySQLdb
import  qdarkstyle
import  urllib2
import  login
import urllib
import json




from PyQt4 import QtCore, QtGui, uic
from PyQt4 import QtCore, QtGui, uic

qtCreatorFile = "questions.ui"

Ui_MainWindow, QtBaseClass = uic.loadUiType(qtCreatorFile)
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
                label.setMaximumSize(700, 300)
                label.setStyleSheet('color:#ea4c88;font-size:14px;font-family:Lato; padding-left:5px;text-align:left')
                label.setText(
                    "\t\t\t\t\t\t<b><br><br>&nbsp;&nbsp;&nbsp;Instructions:</b><br><br>1.You can answer all the questions any number of times until you click submit button <br>or the test time expires.<br><br>\n\t2.You can change your option at any point of time.<br><br>\n\t3.Once you click the submit button you will not be allowed to change your options.<br><br>\n\n\tDo well. <b>ALL THE BEST!!!</b>")
        def doit1(self):
            self.setWindowTitle("Instructions")
            self.disp1()
            self.setStyleSheet(qdarkstyle.load_stylesheet(pyside=False))
            self.setGeometry(QtCore.QRect(100, 100, 700, 500))
            self.center()
            self.show()

pop=MyPopup1()
pop.doit1()
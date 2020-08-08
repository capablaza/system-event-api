package main

import (
	"testing"
)

func TestTranslateDateJan(t *testing.T) {
	inputDate := "06 Jan 2020 17:20:47"
	dateExpected := "06 Ene 2020 17:20:47"
	dateTranslated := translateDate(inputDate)

	if dateTranslated != dateExpected {
		t.Errorf("The result expected is :  %s , but the response was : %s", dateExpected, dateTranslated)
	}
}

func TestTranslateDateFeb(t *testing.T) {
	inputDate := "06 Feb 2020 17:20:47"
	dateExpected := "06 Feb 2020 17:20:47"
	dateTranslated := translateDate(inputDate)

	if dateTranslated != dateExpected {
		t.Errorf("The result expected is :  %s , but the response was : %s", dateExpected, dateTranslated)
	}
}

func TestTranslateDateMar(t *testing.T) {
	inputDate := "06 Mar 2020 17:20:47"
	dateExpected := "06 Mar 2020 17:20:47"
	dateTranslated := translateDate(inputDate)

	if dateTranslated != dateExpected {
		t.Errorf("The result expected is :  %s , but the response was : %s", dateExpected, dateTranslated)
	}
}

func TestTranslateDateApr(t *testing.T) {
	inputDate := "06 Apr 2020 17:20:47"
	dateExpected := "06 Abr 2020 17:20:47"
	dateTranslated := translateDate(inputDate)

	if dateTranslated != dateExpected {
		t.Errorf("The result expected is :  %s , but the response was : %s", dateExpected, dateTranslated)
	}
}

func TestTranslateDateMay(t *testing.T) {
	inputDate := "06 May 2020 17:20:47"
	dateExpected := "06 May 2020 17:20:47"
	dateTranslated := translateDate(inputDate)

	if dateTranslated != dateExpected {
		t.Errorf("The result expected is :  %s , but the response was : %s", dateExpected, dateTranslated)
	}
}

func TestTranslateDateJun(t *testing.T) {
	inputDate := "06 Jun 2020 17:20:47"
	dateExpected := "06 Jun 2020 17:20:47"
	dateTranslated := translateDate(inputDate)

	if dateTranslated != dateExpected {
		t.Errorf("The result expected is :  %s , but the response was : %s", dateExpected, dateTranslated)
	}
}
func TestTranslateDateJul(t *testing.T) {
	inputDate := "06 Jul 2020 17:20:47"
	dateExpected := "06 Jul 2020 17:20:47"
	dateTranslated := translateDate(inputDate)

	if dateTranslated != dateExpected {
		t.Errorf("The result expected is :  %s , but the response was : %s", dateExpected, dateTranslated)
	}
}

func TestTranslateDateAgu(t *testing.T) {
	inputDate := "06 Aug 2020 17:20:47"
	dateExpected := "06 Ago 2020 17:20:47"
	dateTranslated := translateDate(inputDate)

	if dateTranslated != dateExpected {
		t.Errorf("The result expected is :  %s , but the response was : %s", dateExpected, dateTranslated)
	}
}

func TestTranslateDateSep(t *testing.T) {
	inputDate := "06 Sep 2020 17:20:47"
	dateExpected := "06 Sep 2020 17:20:47"
	dateTranslated := translateDate(inputDate)

	if dateTranslated != dateExpected {
		t.Errorf("The result expected is :  %s , but the response was : %s", dateExpected, dateTranslated)
	}
}

func TestTranslateDateOct(t *testing.T) {
	inputDate := "06 Oct 2020 17:20:47"
	dateExpected := "06 Oct 2020 17:20:47"
	dateTranslated := translateDate(inputDate)

	if dateTranslated != dateExpected {
		t.Errorf("The result expected is :  %s , but the response was : %s", dateExpected, dateTranslated)
	}
}

func TestTranslateDateNov(t *testing.T) {
	inputDate := "06 Nov 2020 17:20:47"
	dateExpected := "06 Nov 2020 17:20:47"
	dateTranslated := translateDate(inputDate)

	if dateTranslated != dateExpected {
		t.Errorf("The result expected is :  %s , but the response was : %s", dateExpected, dateTranslated)
	}
}

func TestTranslateDateDec(t *testing.T) {
	inputDate := "06 Dec 2020 17:20:47"
	dateExpected := "06 Dic 2020 17:20:47"
	dateTranslated := translateDate(inputDate)

	if dateTranslated != dateExpected {
		t.Errorf("The result expected is :  %s , but the response was : %s", dateExpected, dateTranslated)
	}
}

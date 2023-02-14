package jdt

import "github.com/vlipovetskii/go-zmanim/hebrewcalendar/timeutil/gdt"

const (

	/*
		JewishEpoch the Jewish epoch using the RD (Rata Die/Fixed Date or Reingold Dershowitz) day used in Calendrical Calculations.
		Day 1 is January 1, 0001, Gregorian
	*/
	JewishEpoch gdt.GDay = -1373429

	// ChalakimPerMinute the number  of chalakim (18) in a minute.
	ChalakimPerMinute MoladChalakim = 18

	// ChalakimPerHour the number  of chalakim (1080) in an hour
	ChalakimPerHour MoladChalakim = 1080

	// ChalakimPerDay The number of chalakim (25,920) in a 24 hours day.
	ChalakimPerDay MoladChalakim64 = 25920 // 24 * 1080

	/*
		ChalakimPerMonth the number  of chalakim in an average Jewish month. A month has 29 days, 12 hours and 793
		chalakim (44 minutes and 3.3 seconds) for a total of 765,433 chalakim
	*/
	ChalakimPerMonth MoladChalakim64 = 765433 // (29 * 24 + 12) * 1080 + 793

	/*
		ChalakimMoladTohu days from the beginning of Sunday till molad BaHaRaD.
		Calculated as 1 day, 5 hours and 204 chalakim = (24 + 5) * 1080 + 204 = 31524
	*/
	ChalakimMoladTohu MoladChalakim64 = 31524

	/*
		Chaserim a short year, where both Heshvan and KISLEV are 29 days.
	*/
	Chaserim int32 = 0

	/*
		Kesidran an ordered year, where Heshvan is 29 days and KISLEV is 30 days.
	*/
	Kesidran int32 = 1

	/*
		Shelaimim a long year, where both Heshvan and KISLEV are 30 days.
	*/
	Shelaimim int32 = 2
)

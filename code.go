package main

var code = map[rune]string{
	'A': ".-",
	'B': "-...",
	'C': "-.-.",
	'D': "-..",
	'E': ".",
	'É': "..-..",
	'F': "..-.",
	'G': "--.",
	'H': "....",
	'I': "..",
	'J': ".---",
	'K': "-.-", // Invitation to transmit
	'L': ".-..",
	'M': "--",
	'N': "-.",
	'O': "---",
	'P': ".--.",
	'Q': "--.-",
	'R': ".-.",
	'S': "...",
	'T': "-",
	'U': "..-",
	'V': "...-",
	'W': ".--",
	'X': "-..-", // Multiplication sign
	'Y': "-.--",
	'Z': "--..",

	'1': ".----",
	'2': "..---",
	'3': "...--",
	'4': "....-",
	'5': ".....",
	'6': "-....",
	'7': "--...",
	'8': "---..",
	'9': "----.",
	'0': "-----",

	'Ä': ".-.-",
	'Ö': "---.",
	'Ü': "..--",
	'Č': "----",

	'.':  ".-.-.-", // Full stop (period)
	',':  "--..--", // Comma
	':':  "---...", // Colon or division sign
	'?':  "..--..", // Question mark
	'\'': ".----.", // Apostrophe
	'-':  "-....-", // Hyphen or dash or substraction sign
	'/':  "-..-.",  // Fraction bar or division sign
	'(':  "-.--.",  // Left-hand bracket
	')':  "-.--.-", // Right-hand bracket
	'"':  ".-..-.", // Inverted commas
	'=':  "-...-",  // Double hyphen
	'+':  ".-.-.",  // Cross or addition sign
	'@':  ".--.-.", // Commercial at

	// 'Roger': "...-.", // Understood / SN
	// 'Error': "........", // Error / HH Correction
	// 'Wait': ".-..." // Wait / AS
	// 'EOW': "...-.-" // End of work / SK Out
}

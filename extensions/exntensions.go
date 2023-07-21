package extensions

// Extensions is a bitmask of enabled parser extensions.
type Extensions int

// Bit flags representing markdown parsing extensions.
// Use | (or) to specify multiple extensions.
const (
	NoExtensions           Extensions = 0
	NoIntraEmphasis        Extensions = 1 << iota // Ignore emphasis markers inside words
	Tables                                        // Parse tables
	FencedCode                                    // Parse fenced code blocks
	Autolink                                      // Detect embedded URLs that are not explicitly marked
	Strikethrough                                 // Strikethrough text using ~~test~~
	LaxHTMLBlocks                                 // Loosen up HTML block parsing rules
	SpaceHeadings                                 // Be strict about prefix heading rules
	HardLineBreak                                 // Translate newlines into line breaks
	NonBlockingSpace                              // Translate backspace spaces into line non-blocking spaces
	TabSizeEight                                  // Expand tabs to eight spaces instead of four
	Footnotes                                     // Pandoc-style footnotes
	NoEmptyLineBeforeBlock                        // No need to insert an empty line to start a (code, quote, ordered list, unordered list) block
	HeadingIDs                                    // specify heading IDs  with {#id}
	Titleblock                                    // Titleblock ala pandoc
	AutoHeadingIDs                                // Create the heading ID from the text
	BackslashLineBreak                            // Translate trailing backslashes into line breaks
	DefinitionLists                               // Parse definition lists
	MathJax                                       // Parse MathJax
	OrderedListStart                              // Keep track of the first number used when starting an ordered list.
	Attributes                                    // Block Attributes
	SuperSubscript                                // Super- and subscript support: 2^10^, H~2~O.
	EmptyLinesBreakList                           // 2 empty lines break out of list
	Includes                                      // Support including other files.
	Mmark                                         // Support Mmark syntax, see https://mmark.miek.nl/post/syntax/

	CommonExtensions Extensions = NoIntraEmphasis | Tables | FencedCode |
		Autolink | Strikethrough | SpaceHeadings | HeadingIDs |
		BackslashLineBreak | DefinitionLists | MathJax
)

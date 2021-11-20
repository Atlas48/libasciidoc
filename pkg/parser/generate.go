package parser

//go:generate pigeon -optimize-parser -optimize-grammar -alternate-entrypoints DocumentFragment,PostReplacementsGroup,DelimitedBlockElements,ElementAttributesGroup,DocumentFragmentWithinVerbatimBlock,NormalGroup,HeaderGroup,AttributesGroup,MacrosGroup,QuotesGroup,NoneGroup,ReplacementsGroup,SpecialCharactersGroup,VerbatimGroup,FileLocation,IncludedFileLine,MarkdownQuoteAttribution,BlockAttributes,InlineAttributes,TableColumnsAttribute,LineRanges,TagRanges -o parser.go parser.peg

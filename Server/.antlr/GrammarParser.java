// Generated from c:\Users\50246\Desktop\Ingeniería-Sexto Semestre\Compiladores 2\Proyecto 1\Server\Grammar.g4 by ANTLR 4.9.2

    import "Server/Interfaces"
    import "Server/Environment"
    import "Server/Expression"
    import "Server/Instruction"
    import "strings"

import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class GrammarParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		STRINGS=1, INTS=2, FLOATS=3, BOOLS=4, CHARACTERS=5, VAR=6, NIL=7, TRUE=8, 
		FALSE=9, PRINT=10, IF=11, ELSE=12, WHILE=13, LET=14, FOR=15, IN=16, NUMBER=17, 
		STRING=18, ID=19, DIF=20, IG_IG=21, NOT=22, OR=23, AND=24, IG=25, MAY_IG=26, 
		MEN_IG=27, MAYOR=28, MENOR=29, MUL=30, DIV=31, ADD=32, SUB=33, PERCENT=34, 
		PARIZQ=35, PARDER=36, LLAVEIZQ=37, LLAVEDER=38, DOUBLEPTS=39, QUESTION=40, 
		WHITESPACE=41, COMMENT=42, LINE_COMMENT=43;
	public static final int
		RULE_s = 0, RULE_block = 1, RULE_instruction = 2, RULE_printstmt = 3, 
		RULE_declarestmt = 4, RULE_constantstmt = 5, RULE_ifstmt = 6, RULE_whilestmt = 7, 
		RULE_forstmt = 8, RULE_expr = 9, RULE_accessstmt = 10, RULE_type = 11;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "block", "instruction", "printstmt", "declarestmt", "constantstmt", 
			"ifstmt", "whilestmt", "forstmt", "expr", "accessstmt", "type"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'String'", "'Int'", "'Float'", "'Bool'", "'Character'", "'var'", 
			"'nil'", "'true'", "'false'", "'print'", "'if'", "'else'", "'while'", 
			"'let'", "'for'", "'in'", null, null, null, "'!='", "'=='", "'!'", "'||'", 
			"'&&'", "'='", "'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'-'", 
			"'%'", "'('", "')'", "'{'", "'}'", "':'", "'?'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "STRINGS", "INTS", "FLOATS", "BOOLS", "CHARACTERS", "VAR", "NIL", 
			"TRUE", "FALSE", "PRINT", "IF", "ELSE", "WHILE", "LET", "FOR", "IN", 
			"NUMBER", "STRING", "ID", "DIF", "IG_IG", "NOT", "OR", "AND", "IG", "MAY_IG", 
			"MEN_IG", "MAYOR", "MENOR", "MUL", "DIV", "ADD", "SUB", "PERCENT", "PARIZQ", 
			"PARDER", "LLAVEIZQ", "LLAVEDER", "DOUBLEPTS", "QUESTION", "WHITESPACE", 
			"COMMENT", "LINE_COMMENT"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Grammar.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public GrammarParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class SContext extends ParserRuleContext {
		public []interface{} code;
		public BlockContext block;
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode EOF() { return getToken(GrammarParser.EOF, 0); }
		public SContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_s; }
	}

	public final SContext s() throws RecognitionException {
		SContext _localctx = new SContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_s);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(24);
			((SContext)_localctx).block = block();
			setState(25);
			match(EOF);
			_localctx.code = ((SContext)_localctx).block.blk
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class BlockContext extends ParserRuleContext {
		public []interface{} blk;
		public InstructionContext instruction;
		public List<InstructionContext> ins = new ArrayList<InstructionContext>();
		public List<InstructionContext> instruction() {
			return getRuleContexts(InstructionContext.class);
		}
		public InstructionContext instruction(int i) {
			return getRuleContext(InstructionContext.class,i);
		}
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_block);

		_localctx.blk = []interface{}{}
		    var listInt []IInstructionContext
		  
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(29); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(28);
				((BlockContext)_localctx).instruction = instruction();
				((BlockContext)_localctx).ins.add(((BlockContext)_localctx).instruction);
				}
				}
				setState(31); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << VAR) | (1L << PRINT) | (1L << IF))) != 0) );

			        listInt = localctx.(*BlockContext).GetIns()
			        for _, e := range listInt {
			            _localctx.blk = append(_localctx.blk, e.GetInst())
			        }
			    
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class InstructionContext extends ParserRuleContext {
		public interfaces.Instruction inst;
		public PrintstmtContext printstmt;
		public DeclarestmtContext declarestmt;
		public IfstmtContext ifstmt;
		public PrintstmtContext printstmt() {
			return getRuleContext(PrintstmtContext.class,0);
		}
		public DeclarestmtContext declarestmt() {
			return getRuleContext(DeclarestmtContext.class,0);
		}
		public IfstmtContext ifstmt() {
			return getRuleContext(IfstmtContext.class,0);
		}
		public InstructionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruction; }
	}

	public final InstructionContext instruction() throws RecognitionException {
		InstructionContext _localctx = new InstructionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_instruction);
		try {
			setState(44);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PRINT:
				enterOuterAlt(_localctx, 1);
				{
				setState(35);
				((InstructionContext)_localctx).printstmt = printstmt();
				 _localctx.inst = ((InstructionContext)_localctx).printstmt.prnt
				}
				break;
			case VAR:
				enterOuterAlt(_localctx, 2);
				{
				setState(38);
				((InstructionContext)_localctx).declarestmt = declarestmt();
				_localctx.inst = ((InstructionContext)_localctx).declarestmt.dec
				}
				break;
			case IF:
				enterOuterAlt(_localctx, 3);
				{
				setState(41);
				((InstructionContext)_localctx).ifstmt = ifstmt();
				 _localctx.inst = ((InstructionContext)_localctx).ifstmt.ift 
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PrintstmtContext extends ParserRuleContext {
		public interfaces.Instruction prnt;
		public Token PRINT;
		public ExprContext expr;
		public TerminalNode PRINT() { return getToken(GrammarParser.PRINT, 0); }
		public TerminalNode PARIZQ() { return getToken(GrammarParser.PARIZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(GrammarParser.PARDER, 0); }
		public PrintstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_printstmt; }
	}

	public final PrintstmtContext printstmt() throws RecognitionException {
		PrintstmtContext _localctx = new PrintstmtContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_printstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(46);
			((PrintstmtContext)_localctx).PRINT = match(PRINT);
			setState(47);
			match(PARIZQ);
			setState(48);
			((PrintstmtContext)_localctx).expr = expr(0);
			setState(49);
			match(PARDER);
			 _localctx.prnt = instructions.NewPrint((((PrintstmtContext)_localctx).PRINT!=null?((PrintstmtContext)_localctx).PRINT.getLine():0), (((PrintstmtContext)_localctx).PRINT!=null?((PrintstmtContext)_localctx).PRINT.getCharPositionInLine():0), ((PrintstmtContext)_localctx).expr.e)
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DeclarestmtContext extends ParserRuleContext {
		public interfaces.Instruction dec;
		public Token VAR;
		public Token ID;
		public TypeContext type;
		public ExprContext expr;
		public TerminalNode VAR() { return getToken(GrammarParser.VAR, 0); }
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode DOUBLEPTS() { return getToken(GrammarParser.DOUBLEPTS, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode IG() { return getToken(GrammarParser.IG, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode QUESTION() { return getToken(GrammarParser.QUESTION, 0); }
		public DeclarestmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declarestmt; }
	}

	public final DeclarestmtContext declarestmt() throws RecognitionException {
		DeclarestmtContext _localctx = new DeclarestmtContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_declarestmt);
		try {
			setState(70);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(52);
				((DeclarestmtContext)_localctx).VAR = match(VAR);
				setState(53);
				((DeclarestmtContext)_localctx).ID = match(ID);
				setState(54);
				match(DOUBLEPTS);
				setState(55);
				((DeclarestmtContext)_localctx).type = type();
				setState(56);
				match(IG);
				setState(57);
				((DeclarestmtContext)_localctx).expr = expr(0);
				_localctx.dec = instructions.NewTodeclare((((DeclarestmtContext)_localctx).VAR!=null?((DeclarestmtContext)_localctx).VAR.getLine():0), (((DeclarestmtContext)_localctx).VAR!=null?((DeclarestmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarestmtContext)_localctx).ID!=null?((DeclarestmtContext)_localctx).ID.getText():null), ((DeclarestmtContext)_localctx).type.t, ((DeclarestmtContext)_localctx).expr.e, false)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(60);
				match(VAR);
				setState(61);
				match(ID);
				setState(62);
				match(IG);
				setState(63);
				expr(0);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(64);
				match(VAR);
				setState(65);
				match(ID);
				setState(66);
				match(DOUBLEPTS);
				setState(67);
				type();
				setState(68);
				match(QUESTION);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ConstantstmtContext extends ParserRuleContext {
		public TerminalNode LET() { return getToken(GrammarParser.LET, 0); }
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode DOUBLEPTS() { return getToken(GrammarParser.DOUBLEPTS, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode IG() { return getToken(GrammarParser.IG, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ConstantstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_constantstmt; }
	}

	public final ConstantstmtContext constantstmt() throws RecognitionException {
		ConstantstmtContext _localctx = new ConstantstmtContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_constantstmt);
		try {
			setState(83);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(72);
				match(LET);
				setState(73);
				match(ID);
				setState(74);
				match(DOUBLEPTS);
				setState(75);
				type();
				setState(76);
				match(IG);
				setState(77);
				expr(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(79);
				match(LET);
				setState(80);
				match(ID);
				setState(81);
				match(IG);
				setState(82);
				expr(0);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class IfstmtContext extends ParserRuleContext {
		public interfaces.Instruction ift;
		public ExprContext left;
		public TerminalNode IF() { return getToken(GrammarParser.IF, 0); }
		public TerminalNode PARIZQ() { return getToken(GrammarParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(GrammarParser.PARDER, 0); }
		public List<TerminalNode> LLAVEIZQ() { return getTokens(GrammarParser.LLAVEIZQ); }
		public TerminalNode LLAVEIZQ(int i) {
			return getToken(GrammarParser.LLAVEIZQ, i);
		}
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public List<TerminalNode> LLAVEDER() { return getTokens(GrammarParser.LLAVEDER); }
		public TerminalNode LLAVEDER(int i) {
			return getToken(GrammarParser.LLAVEDER, i);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode ELSE() { return getToken(GrammarParser.ELSE, 0); }
		public IfstmtContext ifstmt() {
			return getRuleContext(IfstmtContext.class,0);
		}
		public IfstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifstmt; }
	}

	public final IfstmtContext ifstmt() throws RecognitionException {
		IfstmtContext _localctx = new IfstmtContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_ifstmt);
		try {
			setState(115);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(85);
				match(IF);
				setState(86);
				match(PARIZQ);
				setState(87);
				((IfstmtContext)_localctx).left = expr(0);
				setState(88);
				match(PARDER);
				setState(89);
				match(LLAVEIZQ);
				setState(90);
				block();
				setState(91);
				match(LLAVEDER);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(93);
				match(IF);
				setState(94);
				match(PARIZQ);
				setState(95);
				expr(0);
				setState(96);
				match(PARDER);
				setState(97);
				match(LLAVEIZQ);
				setState(98);
				block();
				setState(99);
				match(LLAVEDER);
				setState(100);
				match(ELSE);
				setState(101);
				match(LLAVEIZQ);
				setState(102);
				block();
				setState(103);
				match(LLAVEDER);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(105);
				match(IF);
				setState(106);
				match(PARIZQ);
				setState(107);
				expr(0);
				setState(108);
				match(PARDER);
				setState(109);
				match(LLAVEIZQ);
				setState(110);
				block();
				setState(111);
				match(LLAVEDER);
				setState(112);
				match(ELSE);
				setState(113);
				ifstmt();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class WhilestmtContext extends ParserRuleContext {
		public interfaces.Instruction while;
		public TerminalNode WHILE() { return getToken(GrammarParser.WHILE, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(GrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(GrammarParser.LLAVEDER, 0); }
		public WhilestmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_whilestmt; }
	}

	public final WhilestmtContext whilestmt() throws RecognitionException {
		WhilestmtContext _localctx = new WhilestmtContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_whilestmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(117);
			match(WHILE);
			setState(118);
			expr(0);
			setState(119);
			match(LLAVEIZQ);
			setState(120);
			block();
			setState(121);
			match(LLAVEDER);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ForstmtContext extends ParserRuleContext {
		public interfaces.Instruction for;
		public TerminalNode FOR() { return getToken(GrammarParser.FOR, 0); }
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode IN() { return getToken(GrammarParser.IN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(GrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(GrammarParser.LLAVEDER, 0); }
		public ForstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_forstmt; }
	}

	public final ForstmtContext forstmt() throws RecognitionException {
		ForstmtContext _localctx = new ForstmtContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_forstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(123);
			match(FOR);
			setState(124);
			match(ID);
			setState(125);
			match(IN);
			setState(126);
			expr(0);
			setState(127);
			match(LLAVEIZQ);
			setState(128);
			block();
			setState(129);
			match(LLAVEDER);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExprContext extends ParserRuleContext {
		public interfaces.Expression e;
		public ExprContext left;
		public Token op;
		public ExprContext expr;
		public Token NUMBER;
		public Token STRING;
		public Token TRUE;
		public Token FALSE;
		public AccessstmtContext accessstmt;
		public ExprContext right;
		public TerminalNode NOT() { return getToken(GrammarParser.NOT, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode PARIZQ() { return getToken(GrammarParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(GrammarParser.PARDER, 0); }
		public TerminalNode NUMBER() { return getToken(GrammarParser.NUMBER, 0); }
		public TerminalNode STRING() { return getToken(GrammarParser.STRING, 0); }
		public TerminalNode TRUE() { return getToken(GrammarParser.TRUE, 0); }
		public TerminalNode FALSE() { return getToken(GrammarParser.FALSE, 0); }
		public AccessstmtContext accessstmt() {
			return getRuleContext(AccessstmtContext.class,0);
		}
		public TerminalNode MUL() { return getToken(GrammarParser.MUL, 0); }
		public TerminalNode DIV() { return getToken(GrammarParser.DIV, 0); }
		public TerminalNode ADD() { return getToken(GrammarParser.ADD, 0); }
		public TerminalNode SUB() { return getToken(GrammarParser.SUB, 0); }
		public TerminalNode PERCENT() { return getToken(GrammarParser.PERCENT, 0); }
		public TerminalNode MAY_IG() { return getToken(GrammarParser.MAY_IG, 0); }
		public TerminalNode MAYOR() { return getToken(GrammarParser.MAYOR, 0); }
		public TerminalNode MEN_IG() { return getToken(GrammarParser.MEN_IG, 0); }
		public TerminalNode MENOR() { return getToken(GrammarParser.MENOR, 0); }
		public TerminalNode IG_IG() { return getToken(GrammarParser.IG_IG, 0); }
		public TerminalNode DIF() { return getToken(GrammarParser.DIF, 0); }
		public TerminalNode AND() { return getToken(GrammarParser.AND, 0); }
		public TerminalNode OR() { return getToken(GrammarParser.OR, 0); }
		public ExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr; }
	}

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 18;
		enterRecursionRule(_localctx, 18, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(152);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NOT:
				{
				setState(132);
				((ExprContext)_localctx).op = match(NOT);
				setState(133);
				((ExprContext)_localctx).left = ((ExprContext)_localctx).expr = expr(7);
				 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).left.e) 
				}
				break;
			case PARIZQ:
				{
				setState(136);
				match(PARIZQ);
				setState(137);
				((ExprContext)_localctx).expr = expr(0);
				setState(138);
				match(PARDER);
				 _localctx.e = ((ExprContext)_localctx).expr.e 
				}
				break;
			case NUMBER:
				{
				setState(141);
				((ExprContext)_localctx).NUMBER = match(NUMBER);

				        if (strings.Contains((((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getText():null),".")){
				            num,err := strconv.ParseFloat((((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getText():null), 64);
				            if err!= nil{
				                fmt.Println(err)
				            }
				            _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getLine():0),(((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getCharPositionInLine():0),num,environment.FLOAT)
				        }else{
				            num,err := strconv.Atoi((((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getText():null))
				            if err!= nil{
				                fmt.Println(err)
				            }
				            _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getLine():0),(((ExprContext)_localctx).NUMBER!=null?((ExprContext)_localctx).NUMBER.getCharPositionInLine():0),num,environment.INTEGER)
				        }
				    
				}
				break;
			case STRING:
				{
				setState(143);
				((ExprContext)_localctx).STRING = match(STRING);

				        str := (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getText():null)
				        _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getLine():0), (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getCharPositionInLine():0), str[1:len(str)-1],environment.STRING)
				    
				}
				break;
			case TRUE:
				{
				setState(145);
				((ExprContext)_localctx).TRUE = match(TRUE);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).TRUE!=null?((ExprContext)_localctx).TRUE.getLine():0), (((ExprContext)_localctx).TRUE!=null?((ExprContext)_localctx).TRUE.getCharPositionInLine():0), true, environment.BOOLEAN) 
				}
				break;
			case FALSE:
				{
				setState(147);
				((ExprContext)_localctx).FALSE = match(FALSE);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).FALSE!=null?((ExprContext)_localctx).FALSE.getLine():0), (((ExprContext)_localctx).FALSE!=null?((ExprContext)_localctx).FALSE.getCharPositionInLine():0), false, environment.BOOLEAN) 
				}
				break;
			case ID:
				{
				setState(149);
				((ExprContext)_localctx).accessstmt = accessstmt();
				_localctx.e = ((ExprContext)_localctx).accessstmt.access
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(191);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(189);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
					case 1:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(154);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(155);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==MUL || _la==DIV) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(156);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(15);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 2:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(159);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(160);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << ADD) | (1L << SUB) | (1L << PERCENT))) != 0)) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(161);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(14);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 3:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(164);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(165);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==MAY_IG || _la==MAYOR) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(166);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(13);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 4:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(169);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(170);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==MEN_IG || _la==MENOR) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(171);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(12);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 5:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(174);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(175);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==DIF || _la==IG_IG) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(176);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(11);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 6:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(179);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(180);
						((ExprContext)_localctx).op = match(AND);
						setState(181);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(10);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 7:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(184);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(185);
						((ExprContext)_localctx).op = match(OR);
						setState(186);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(9);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					}
					} 
				}
				setState(193);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class AccessstmtContext extends ParserRuleContext {
		public interfaces.Expression access;
		public Token op;
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public AccessstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_accessstmt; }
	}

	public final AccessstmtContext accessstmt() throws RecognitionException {
		AccessstmtContext _localctx = new AccessstmtContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_accessstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(194);
			((AccessstmtContext)_localctx).op = match(ID);
			_localctx.access = expressions.NewAccess((((AccessstmtContext)_localctx).op!=null?((AccessstmtContext)_localctx).op.getLine():0), (((AccessstmtContext)_localctx).op!=null?((AccessstmtContext)_localctx).op.getCharPositionInLine():0), (((AccessstmtContext)_localctx).op!=null?((AccessstmtContext)_localctx).op.getText():null))
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypeContext extends ParserRuleContext {
		public environment.TipoExpresion t;
		public TerminalNode STRINGS() { return getToken(GrammarParser.STRINGS, 0); }
		public TerminalNode INTS() { return getToken(GrammarParser.INTS, 0); }
		public TerminalNode FLOATS() { return getToken(GrammarParser.FLOATS, 0); }
		public TerminalNode BOOLS() { return getToken(GrammarParser.BOOLS, 0); }
		public TerminalNode CHARACTERS() { return getToken(GrammarParser.CHARACTERS, 0); }
		public TypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type; }
	}

	public final TypeContext type() throws RecognitionException {
		TypeContext _localctx = new TypeContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_type);
		try {
			setState(207);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRINGS:
				enterOuterAlt(_localctx, 1);
				{
				setState(197);
				match(STRINGS);
				_localctx.t = environment.STRING
				}
				break;
			case INTS:
				enterOuterAlt(_localctx, 2);
				{
				setState(199);
				match(INTS);
				_localctx.t = environment.INTEGER
				}
				break;
			case FLOATS:
				enterOuterAlt(_localctx, 3);
				{
				setState(201);
				match(FLOATS);
				_localctx.t = environment.FLOAT
				}
				break;
			case BOOLS:
				enterOuterAlt(_localctx, 4);
				{
				setState(203);
				match(BOOLS);
				_localctx.t = environment.BOOLEAN 
				}
				break;
			case CHARACTERS:
				enterOuterAlt(_localctx, 5);
				{
				setState(205);
				match(CHARACTERS);
				_localctx.t = environment.CHARACTER
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 9:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 14);
		case 1:
			return precpred(_ctx, 13);
		case 2:
			return precpred(_ctx, 12);
		case 3:
			return precpred(_ctx, 11);
		case 4:
			return precpred(_ctx, 10);
		case 5:
			return precpred(_ctx, 9);
		case 6:
			return precpred(_ctx, 8);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3-\u00d4\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\3\2\3\2\3\2\3\2\3\3\6\3 \n\3\r\3\16\3!\3\3\3\3\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4/\n\4\3\5\3\5\3\5\3\5\3\5\3\5\3\6"+
		"\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5"+
		"\6I\n\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\5\7V\n\7\3\b\3\b\3"+
		"\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b"+
		"\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\5\bv\n\b\3\t\3\t\3\t\3\t\3\t"+
		"\3\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\5\13\u009b\n\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\7\13\u00c0\n\13\f\13"+
		"\16\13\u00c3\13\13\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r"+
		"\5\r\u00d2\n\r\3\r\2\3\24\16\2\4\6\b\n\f\16\20\22\24\26\30\2\7\3\2 !\3"+
		"\2\"$\4\2\34\34\36\36\4\2\35\35\37\37\3\2\26\27\2\u00e0\2\32\3\2\2\2\4"+
		"\37\3\2\2\2\6.\3\2\2\2\b\60\3\2\2\2\nH\3\2\2\2\fU\3\2\2\2\16u\3\2\2\2"+
		"\20w\3\2\2\2\22}\3\2\2\2\24\u009a\3\2\2\2\26\u00c4\3\2\2\2\30\u00d1\3"+
		"\2\2\2\32\33\5\4\3\2\33\34\7\2\2\3\34\35\b\2\1\2\35\3\3\2\2\2\36 \5\6"+
		"\4\2\37\36\3\2\2\2 !\3\2\2\2!\37\3\2\2\2!\"\3\2\2\2\"#\3\2\2\2#$\b\3\1"+
		"\2$\5\3\2\2\2%&\5\b\5\2&\'\b\4\1\2\'/\3\2\2\2()\5\n\6\2)*\b\4\1\2*/\3"+
		"\2\2\2+,\5\16\b\2,-\b\4\1\2-/\3\2\2\2.%\3\2\2\2.(\3\2\2\2.+\3\2\2\2/\7"+
		"\3\2\2\2\60\61\7\f\2\2\61\62\7%\2\2\62\63\5\24\13\2\63\64\7&\2\2\64\65"+
		"\b\5\1\2\65\t\3\2\2\2\66\67\7\b\2\2\678\7\25\2\289\7)\2\29:\5\30\r\2:"+
		";\7\33\2\2;<\5\24\13\2<=\b\6\1\2=I\3\2\2\2>?\7\b\2\2?@\7\25\2\2@A\7\33"+
		"\2\2AI\5\24\13\2BC\7\b\2\2CD\7\25\2\2DE\7)\2\2EF\5\30\r\2FG\7*\2\2GI\3"+
		"\2\2\2H\66\3\2\2\2H>\3\2\2\2HB\3\2\2\2I\13\3\2\2\2JK\7\20\2\2KL\7\25\2"+
		"\2LM\7)\2\2MN\5\30\r\2NO\7\33\2\2OP\5\24\13\2PV\3\2\2\2QR\7\20\2\2RS\7"+
		"\25\2\2ST\7\33\2\2TV\5\24\13\2UJ\3\2\2\2UQ\3\2\2\2V\r\3\2\2\2WX\7\r\2"+
		"\2XY\7%\2\2YZ\5\24\13\2Z[\7&\2\2[\\\7\'\2\2\\]\5\4\3\2]^\7(\2\2^v\3\2"+
		"\2\2_`\7\r\2\2`a\7%\2\2ab\5\24\13\2bc\7&\2\2cd\7\'\2\2de\5\4\3\2ef\7("+
		"\2\2fg\7\16\2\2gh\7\'\2\2hi\5\4\3\2ij\7(\2\2jv\3\2\2\2kl\7\r\2\2lm\7%"+
		"\2\2mn\5\24\13\2no\7&\2\2op\7\'\2\2pq\5\4\3\2qr\7(\2\2rs\7\16\2\2st\5"+
		"\16\b\2tv\3\2\2\2uW\3\2\2\2u_\3\2\2\2uk\3\2\2\2v\17\3\2\2\2wx\7\17\2\2"+
		"xy\5\24\13\2yz\7\'\2\2z{\5\4\3\2{|\7(\2\2|\21\3\2\2\2}~\7\21\2\2~\177"+
		"\7\25\2\2\177\u0080\7\22\2\2\u0080\u0081\5\24\13\2\u0081\u0082\7\'\2\2"+
		"\u0082\u0083\5\4\3\2\u0083\u0084\7(\2\2\u0084\23\3\2\2\2\u0085\u0086\b"+
		"\13\1\2\u0086\u0087\7\30\2\2\u0087\u0088\5\24\13\t\u0088\u0089\b\13\1"+
		"\2\u0089\u009b\3\2\2\2\u008a\u008b\7%\2\2\u008b\u008c\5\24\13\2\u008c"+
		"\u008d\7&\2\2\u008d\u008e\b\13\1\2\u008e\u009b\3\2\2\2\u008f\u0090\7\23"+
		"\2\2\u0090\u009b\b\13\1\2\u0091\u0092\7\24\2\2\u0092\u009b\b\13\1\2\u0093"+
		"\u0094\7\n\2\2\u0094\u009b\b\13\1\2\u0095\u0096\7\13\2\2\u0096\u009b\b"+
		"\13\1\2\u0097\u0098\5\26\f\2\u0098\u0099\b\13\1\2\u0099\u009b\3\2\2\2"+
		"\u009a\u0085\3\2\2\2\u009a\u008a\3\2\2\2\u009a\u008f\3\2\2\2\u009a\u0091"+
		"\3\2\2\2\u009a\u0093\3\2\2\2\u009a\u0095\3\2\2\2\u009a\u0097\3\2\2\2\u009b"+
		"\u00c1\3\2\2\2\u009c\u009d\f\20\2\2\u009d\u009e\t\2\2\2\u009e\u009f\5"+
		"\24\13\21\u009f\u00a0\b\13\1\2\u00a0\u00c0\3\2\2\2\u00a1\u00a2\f\17\2"+
		"\2\u00a2\u00a3\t\3\2\2\u00a3\u00a4\5\24\13\20\u00a4\u00a5\b\13\1\2\u00a5"+
		"\u00c0\3\2\2\2\u00a6\u00a7\f\16\2\2\u00a7\u00a8\t\4\2\2\u00a8\u00a9\5"+
		"\24\13\17\u00a9\u00aa\b\13\1\2\u00aa\u00c0\3\2\2\2\u00ab\u00ac\f\r\2\2"+
		"\u00ac\u00ad\t\5\2\2\u00ad\u00ae\5\24\13\16\u00ae\u00af\b\13\1\2\u00af"+
		"\u00c0\3\2\2\2\u00b0\u00b1\f\f\2\2\u00b1\u00b2\t\6\2\2\u00b2\u00b3\5\24"+
		"\13\r\u00b3\u00b4\b\13\1\2\u00b4\u00c0\3\2\2\2\u00b5\u00b6\f\13\2\2\u00b6"+
		"\u00b7\7\32\2\2\u00b7\u00b8\5\24\13\f\u00b8\u00b9\b\13\1\2\u00b9\u00c0"+
		"\3\2\2\2\u00ba\u00bb\f\n\2\2\u00bb\u00bc\7\31\2\2\u00bc\u00bd\5\24\13"+
		"\13\u00bd\u00be\b\13\1\2\u00be\u00c0\3\2\2\2\u00bf\u009c\3\2\2\2\u00bf"+
		"\u00a1\3\2\2\2\u00bf\u00a6\3\2\2\2\u00bf\u00ab\3\2\2\2\u00bf\u00b0\3\2"+
		"\2\2\u00bf\u00b5\3\2\2\2\u00bf\u00ba\3\2\2\2\u00c0\u00c3\3\2\2\2\u00c1"+
		"\u00bf\3\2\2\2\u00c1\u00c2\3\2\2\2\u00c2\25\3\2\2\2\u00c3\u00c1\3\2\2"+
		"\2\u00c4\u00c5\7\25\2\2\u00c5\u00c6\b\f\1\2\u00c6\27\3\2\2\2\u00c7\u00c8"+
		"\7\3\2\2\u00c8\u00d2\b\r\1\2\u00c9\u00ca\7\4\2\2\u00ca\u00d2\b\r\1\2\u00cb"+
		"\u00cc\7\5\2\2\u00cc\u00d2\b\r\1\2\u00cd\u00ce\7\6\2\2\u00ce\u00d2\b\r"+
		"\1\2\u00cf\u00d0\7\7\2\2\u00d0\u00d2\b\r\1\2\u00d1\u00c7\3\2\2\2\u00d1"+
		"\u00c9\3\2\2\2\u00d1\u00cb\3\2\2\2\u00d1\u00cd\3\2\2\2\u00d1\u00cf\3\2"+
		"\2\2\u00d2\31\3\2\2\2\13!.HUu\u009a\u00bf\u00c1\u00d1";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
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
		FALSE=9, PRINT=10, IF=11, ELSE=12, WHILE=13, LET=14, NUMBER=15, STRING=16, 
		ID=17, DIF=18, IG_IG=19, NOT=20, OR=21, AND=22, IG=23, MAY_IG=24, MEN_IG=25, 
		MAYOR=26, MENOR=27, MUL=28, DIV=29, ADD=30, SUB=31, PERCENT=32, PARIZQ=33, 
		PARDER=34, LLAVEIZQ=35, LLAVEDER=36, DOUBLEPTS=37, QUESTION=38, WHITESPACE=39, 
		COMMENT=40, LINE_COMMENT=41;
	public static final int
		RULE_s = 0, RULE_block = 1, RULE_instruction = 2, RULE_printstmt = 3, 
		RULE_declarestmt = 4, RULE_constantstmt = 5, RULE_ifstmt = 6, RULE_expr = 7, 
		RULE_type = 8;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "block", "instruction", "printstmt", "declarestmt", "constantstmt", 
			"ifstmt", "expr", "type"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'String'", "'Int'", "'Float'", "'Bool'", "'Character'", "'var'", 
			"'nil'", "'true'", "'false'", "'print'", "'if'", "'else'", "'while'", 
			"'let'", null, null, null, "'!='", "'=='", "'!'", "'||'", "'&&'", "'='", 
			"'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'-'", "'%'", "'('", 
			"')'", "'{'", "'}'", "':'", "'?'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "STRINGS", "INTS", "FLOATS", "BOOLS", "CHARACTERS", "VAR", "NIL", 
			"TRUE", "FALSE", "PRINT", "IF", "ELSE", "WHILE", "LET", "NUMBER", "STRING", 
			"ID", "DIF", "IG_IG", "NOT", "OR", "AND", "IG", "MAY_IG", "MEN_IG", "MAYOR", 
			"MENOR", "MUL", "DIV", "ADD", "SUB", "PERCENT", "PARIZQ", "PARDER", "LLAVEIZQ", 
			"LLAVEDER", "DOUBLEPTS", "QUESTION", "WHITESPACE", "COMMENT", "LINE_COMMENT"
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
			setState(18);
			((SContext)_localctx).block = block();
			setState(19);
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
			setState(23); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(22);
				((BlockContext)_localctx).instruction = instruction();
				((BlockContext)_localctx).ins.add(((BlockContext)_localctx).instruction);
				}
				}
				setState(25); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==PRINT || _la==IF );

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
		public PrintstmtContext printstmt() {
			return getRuleContext(PrintstmtContext.class,0);
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
			setState(35);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PRINT:
				enterOuterAlt(_localctx, 1);
				{
				setState(29);
				((InstructionContext)_localctx).printstmt = printstmt();
				 _localctx.inst = ((InstructionContext)_localctx).printstmt.prnt
				}
				break;
			case IF:
				enterOuterAlt(_localctx, 2);
				{
				setState(32);
				ifstmt();
				 
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
			setState(37);
			((PrintstmtContext)_localctx).PRINT = match(PRINT);
			setState(38);
			match(PARIZQ);
			setState(39);
			((PrintstmtContext)_localctx).expr = expr(0);
			setState(40);
			match(PARDER);
			 _localctx.prnt = instructions.NewPrint((((PrintstmtContext)_localctx).PRINT!=null?((PrintstmtContext)_localctx).PRINT.getLine():0),(((PrintstmtContext)_localctx).PRINT!=null?((PrintstmtContext)_localctx).PRINT.getCharPositionInLine():0),((PrintstmtContext)_localctx).expr.e)
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
			setState(60);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(43);
				match(VAR);
				setState(44);
				match(ID);
				setState(45);
				match(DOUBLEPTS);
				setState(46);
				type();
				setState(47);
				match(IG);
				setState(48);
				expr(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(50);
				match(VAR);
				setState(51);
				match(ID);
				setState(52);
				match(IG);
				setState(53);
				expr(0);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(54);
				match(VAR);
				setState(55);
				match(ID);
				setState(56);
				match(DOUBLEPTS);
				setState(57);
				type();
				setState(58);
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
			setState(73);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(62);
				match(LET);
				setState(63);
				match(ID);
				setState(64);
				match(DOUBLEPTS);
				setState(65);
				type();
				setState(66);
				match(IG);
				setState(67);
				expr(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(69);
				match(LET);
				setState(70);
				match(ID);
				setState(71);
				match(IG);
				setState(72);
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
		public TerminalNode IF() { return getToken(GrammarParser.IF, 0); }
		public TerminalNode PARIZQ() { return getToken(GrammarParser.PARIZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
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
			setState(105);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(75);
				match(IF);
				setState(76);
				match(PARIZQ);
				setState(77);
				expr(0);
				setState(78);
				match(PARDER);
				setState(79);
				match(LLAVEIZQ);
				setState(80);
				block();
				setState(81);
				match(LLAVEDER);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(83);
				match(IF);
				setState(84);
				match(PARIZQ);
				setState(85);
				expr(0);
				setState(86);
				match(PARDER);
				setState(87);
				match(LLAVEIZQ);
				setState(88);
				block();
				setState(89);
				match(LLAVEDER);
				setState(90);
				match(ELSE);
				setState(91);
				match(LLAVEIZQ);
				setState(92);
				block();
				setState(93);
				match(LLAVEDER);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(95);
				match(IF);
				setState(96);
				match(PARIZQ);
				setState(97);
				expr(0);
				setState(98);
				match(PARDER);
				setState(99);
				match(LLAVEIZQ);
				setState(100);
				block();
				setState(101);
				match(LLAVEDER);
				setState(102);
				match(ELSE);
				setState(103);
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

	public static class ExprContext extends ParserRuleContext {
		public interfaces.Expression e;
		public ExprContext left;
		public Token op;
		public ExprContext expr;
		public Token NUMBER;
		public Token STRING;
		public Token TRUE;
		public Token FALSE;
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
		int _startState = 14;
		enterRecursionRule(_localctx, 14, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(125);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NOT:
				{
				setState(108);
				((ExprContext)_localctx).op = match(NOT);
				setState(109);
				((ExprContext)_localctx).left = ((ExprContext)_localctx).expr = expr(6);
				 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).left.e) 
				}
				break;
			case PARIZQ:
				{
				setState(112);
				match(PARIZQ);
				setState(113);
				((ExprContext)_localctx).expr = expr(0);
				setState(114);
				match(PARDER);
				 _localctx.e = ((ExprContext)_localctx).expr.e 
				}
				break;
			case NUMBER:
				{
				setState(117);
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
				setState(119);
				((ExprContext)_localctx).STRING = match(STRING);

				        str := (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getText():null)
				        _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getLine():0), (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getCharPositionInLine():0), str[1:len(str)-1],environment.STRING)
				    
				}
				break;
			case TRUE:
				{
				setState(121);
				((ExprContext)_localctx).TRUE = match(TRUE);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).TRUE!=null?((ExprContext)_localctx).TRUE.getLine():0), (((ExprContext)_localctx).TRUE!=null?((ExprContext)_localctx).TRUE.getCharPositionInLine():0), true, environment.BOOLEAN) 
				}
				break;
			case FALSE:
				{
				setState(123);
				((ExprContext)_localctx).FALSE = match(FALSE);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).FALSE!=null?((ExprContext)_localctx).FALSE.getLine():0), (((ExprContext)_localctx).FALSE!=null?((ExprContext)_localctx).FALSE.getCharPositionInLine():0), false, environment.BOOLEAN) 
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(164);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(162);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
					case 1:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(127);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(128);
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
						setState(129);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(14);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 2:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(132);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(133);
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
						setState(134);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(13);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 3:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(137);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(138);
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
						setState(139);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(12);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 4:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(142);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(143);
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
						setState(144);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(11);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 5:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(147);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(148);
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
						setState(149);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(10);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 6:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(152);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(153);
						((ExprContext)_localctx).op = match(AND);
						setState(154);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(9);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 7:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(157);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(158);
						((ExprContext)_localctx).op = match(OR);
						setState(159);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(8);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					}
					} 
				}
				setState(166);
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

	public static class TypeContext extends ParserRuleContext {
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
		enterRule(_localctx, 16, RULE_type);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(167);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << STRINGS) | (1L << INTS) | (1L << FLOATS) | (1L << BOOLS) | (1L << CHARACTERS))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
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

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 7:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 13);
		case 1:
			return precpred(_ctx, 12);
		case 2:
			return precpred(_ctx, 11);
		case 3:
			return precpred(_ctx, 10);
		case 4:
			return precpred(_ctx, 9);
		case 5:
			return precpred(_ctx, 8);
		case 6:
			return precpred(_ctx, 7);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3+\u00ac\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\3\2\3\2"+
		"\3\2\3\2\3\3\6\3\32\n\3\r\3\16\3\33\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\5"+
		"\4&\n\4\3\5\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6?\n\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3"+
		"\7\3\7\3\7\3\7\5\7L\n\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3"+
		"\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b"+
		"\3\b\5\bl\n\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t"+
		"\3\t\3\t\3\t\3\t\5\t\u0080\n\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t"+
		"\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\7\t\u00a5\n\t\f\t\16\t\u00a8\13\t\3\n\3"+
		"\n\3\n\2\3\20\13\2\4\6\b\n\f\16\20\22\2\b\3\2\36\37\3\2 \"\4\2\32\32\34"+
		"\34\4\2\33\33\35\35\3\2\24\25\3\2\3\7\2\u00b5\2\24\3\2\2\2\4\31\3\2\2"+
		"\2\6%\3\2\2\2\b\'\3\2\2\2\n>\3\2\2\2\fK\3\2\2\2\16k\3\2\2\2\20\177\3\2"+
		"\2\2\22\u00a9\3\2\2\2\24\25\5\4\3\2\25\26\7\2\2\3\26\27\b\2\1\2\27\3\3"+
		"\2\2\2\30\32\5\6\4\2\31\30\3\2\2\2\32\33\3\2\2\2\33\31\3\2\2\2\33\34\3"+
		"\2\2\2\34\35\3\2\2\2\35\36\b\3\1\2\36\5\3\2\2\2\37 \5\b\5\2 !\b\4\1\2"+
		"!&\3\2\2\2\"#\5\16\b\2#$\b\4\1\2$&\3\2\2\2%\37\3\2\2\2%\"\3\2\2\2&\7\3"+
		"\2\2\2\'(\7\f\2\2()\7#\2\2)*\5\20\t\2*+\7$\2\2+,\b\5\1\2,\t\3\2\2\2-."+
		"\7\b\2\2./\7\23\2\2/\60\7\'\2\2\60\61\5\22\n\2\61\62\7\31\2\2\62\63\5"+
		"\20\t\2\63?\3\2\2\2\64\65\7\b\2\2\65\66\7\23\2\2\66\67\7\31\2\2\67?\5"+
		"\20\t\289\7\b\2\29:\7\23\2\2:;\7\'\2\2;<\5\22\n\2<=\7(\2\2=?\3\2\2\2>"+
		"-\3\2\2\2>\64\3\2\2\2>8\3\2\2\2?\13\3\2\2\2@A\7\20\2\2AB\7\23\2\2BC\7"+
		"\'\2\2CD\5\22\n\2DE\7\31\2\2EF\5\20\t\2FL\3\2\2\2GH\7\20\2\2HI\7\23\2"+
		"\2IJ\7\31\2\2JL\5\20\t\2K@\3\2\2\2KG\3\2\2\2L\r\3\2\2\2MN\7\r\2\2NO\7"+
		"#\2\2OP\5\20\t\2PQ\7$\2\2QR\7%\2\2RS\5\4\3\2ST\7&\2\2Tl\3\2\2\2UV\7\r"+
		"\2\2VW\7#\2\2WX\5\20\t\2XY\7$\2\2YZ\7%\2\2Z[\5\4\3\2[\\\7&\2\2\\]\7\16"+
		"\2\2]^\7%\2\2^_\5\4\3\2_`\7&\2\2`l\3\2\2\2ab\7\r\2\2bc\7#\2\2cd\5\20\t"+
		"\2de\7$\2\2ef\7%\2\2fg\5\4\3\2gh\7&\2\2hi\7\16\2\2ij\5\16\b\2jl\3\2\2"+
		"\2kM\3\2\2\2kU\3\2\2\2ka\3\2\2\2l\17\3\2\2\2mn\b\t\1\2no\7\26\2\2op\5"+
		"\20\t\bpq\b\t\1\2q\u0080\3\2\2\2rs\7#\2\2st\5\20\t\2tu\7$\2\2uv\b\t\1"+
		"\2v\u0080\3\2\2\2wx\7\21\2\2x\u0080\b\t\1\2yz\7\22\2\2z\u0080\b\t\1\2"+
		"{|\7\n\2\2|\u0080\b\t\1\2}~\7\13\2\2~\u0080\b\t\1\2\177m\3\2\2\2\177r"+
		"\3\2\2\2\177w\3\2\2\2\177y\3\2\2\2\177{\3\2\2\2\177}\3\2\2\2\u0080\u00a6"+
		"\3\2\2\2\u0081\u0082\f\17\2\2\u0082\u0083\t\2\2\2\u0083\u0084\5\20\t\20"+
		"\u0084\u0085\b\t\1\2\u0085\u00a5\3\2\2\2\u0086\u0087\f\16\2\2\u0087\u0088"+
		"\t\3\2\2\u0088\u0089\5\20\t\17\u0089\u008a\b\t\1\2\u008a\u00a5\3\2\2\2"+
		"\u008b\u008c\f\r\2\2\u008c\u008d\t\4\2\2\u008d\u008e\5\20\t\16\u008e\u008f"+
		"\b\t\1\2\u008f\u00a5\3\2\2\2\u0090\u0091\f\f\2\2\u0091\u0092\t\5\2\2\u0092"+
		"\u0093\5\20\t\r\u0093\u0094\b\t\1\2\u0094\u00a5\3\2\2\2\u0095\u0096\f"+
		"\13\2\2\u0096\u0097\t\6\2\2\u0097\u0098\5\20\t\f\u0098\u0099\b\t\1\2\u0099"+
		"\u00a5\3\2\2\2\u009a\u009b\f\n\2\2\u009b\u009c\7\30\2\2\u009c\u009d\5"+
		"\20\t\13\u009d\u009e\b\t\1\2\u009e\u00a5\3\2\2\2\u009f\u00a0\f\t\2\2\u00a0"+
		"\u00a1\7\27\2\2\u00a1\u00a2\5\20\t\n\u00a2\u00a3\b\t\1\2\u00a3\u00a5\3"+
		"\2\2\2\u00a4\u0081\3\2\2\2\u00a4\u0086\3\2\2\2\u00a4\u008b\3\2\2\2\u00a4"+
		"\u0090\3\2\2\2\u00a4\u0095\3\2\2\2\u00a4\u009a\3\2\2\2\u00a4\u009f\3\2"+
		"\2\2\u00a5\u00a8\3\2\2\2\u00a6\u00a4\3\2\2\2\u00a6\u00a7\3\2\2\2\u00a7"+
		"\21\3\2\2\2\u00a8\u00a6\3\2\2\2\u00a9\u00aa\t\7\2\2\u00aa\23\3\2\2\2\n"+
		"\33%>Kk\177\u00a4\u00a6";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
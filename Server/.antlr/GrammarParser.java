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
		FALSE=9, PRINT=10, IF=11, ELSE=12, WHILE=13, LET=14, FOR=15, IN=16, SWITCH=17, 
		CASE=18, DEFAULT=19, GUARD=20, CONTINUE=21, RETURN=22, BREAK=23, APPEND=24, 
		REMOVELAST=25, REMOVE=26, AT=27, ISEMPTY=28, COUNT=29, NUMBER=30, STRING=31, 
		ID=32, DIF=33, IG_IG=34, NOT=35, OR=36, AND=37, IG=38, IG_ADD=39, IG_SUB=40, 
		MAY_IG=41, MEN_IG=42, MAYOR=43, MENOR=44, MUL=45, DIV=46, ADD=47, SUB=48, 
		PERCENT=49, PARIZQ=50, PARDER=51, LLAVEIZQ=52, LLAVEDER=53, DOUBLEPTS=54, 
		QUESTION=55, POINT=56, COMA=57, PTCOMA=58, CORCHETEIZQ=59, CORCHETEDER=60, 
		WHITESPACE=61, COMMENT=62, LINE_COMMENT=63;
	public static final int
		RULE_s = 0, RULE_block = 1, RULE_instruction = 2, RULE_printstmt = 3, 
		RULE_declarestmt = 4, RULE_constantstmt = 5, RULE_blockelifs = 6, RULE_ifstmt = 7, 
		RULE_switchstmt = 8, RULE_casestmt = 9, RULE_blockcases = 10, RULE_assignationstmt = 11, 
		RULE_whilestmt = 12, RULE_forstmt = 13, RULE_guardstmt = 14, RULE_transferstmt = 15, 
		RULE_declarevectorstmt = 16, RULE_appendstmt = 17, RULE_removelaststmt = 18, 
		RULE_removestmt = 19, RULE_emptvecstmt = 20, RULE_countvecstmt = 21, RULE_listParams = 22, 
		RULE_expr = 23, RULE_accessstmt = 24, RULE_increaseanddecreasestmt = 25, 
		RULE_type = 26;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "block", "instruction", "printstmt", "declarestmt", "constantstmt", 
			"blockelifs", "ifstmt", "switchstmt", "casestmt", "blockcases", "assignationstmt", 
			"whilestmt", "forstmt", "guardstmt", "transferstmt", "declarevectorstmt", 
			"appendstmt", "removelaststmt", "removestmt", "emptvecstmt", "countvecstmt", 
			"listParams", "expr", "accessstmt", "increaseanddecreasestmt", "type"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'String'", "'Int'", "'Float'", "'Bool'", "'Character'", "'var'", 
			"'nil'", "'true'", "'false'", "'print'", "'if'", "'else'", "'while'", 
			"'let'", "'for'", "'in'", "'switch'", "'case'", "'default'", "'guard'", 
			"'continue'", "'return'", "'break'", "'append'", "'removeLast'", "'remove'", 
			"'at'", "'IsEmpty'", "'count'", null, null, null, "'!='", "'=='", "'!'", 
			"'||'", "'&&'", "'='", "'+='", "'-='", "'>='", "'<='", "'>'", "'<'", 
			"'*'", "'/'", "'+'", "'-'", "'%'", "'('", "')'", "'{'", "'}'", "':'", 
			"'?'", "'.'", "','", "';'", "'['", "']'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "STRINGS", "INTS", "FLOATS", "BOOLS", "CHARACTERS", "VAR", "NIL", 
			"TRUE", "FALSE", "PRINT", "IF", "ELSE", "WHILE", "LET", "FOR", "IN", 
			"SWITCH", "CASE", "DEFAULT", "GUARD", "CONTINUE", "RETURN", "BREAK", 
			"APPEND", "REMOVELAST", "REMOVE", "AT", "ISEMPTY", "COUNT", "NUMBER", 
			"STRING", "ID", "DIF", "IG_IG", "NOT", "OR", "AND", "IG", "IG_ADD", "IG_SUB", 
			"MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", "DIV", "ADD", "SUB", "PERCENT", 
			"PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "DOUBLEPTS", "QUESTION", 
			"POINT", "COMA", "PTCOMA", "CORCHETEIZQ", "CORCHETEDER", "WHITESPACE", 
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
			setState(54);
			((SContext)_localctx).block = block();
			setState(55);
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
		  
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(59); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(58);
					((BlockContext)_localctx).instruction = instruction();
					((BlockContext)_localctx).ins.add(((BlockContext)_localctx).instruction);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(61); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );

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
		public ConstantstmtContext constantstmt;
		public AssignationstmtContext assignationstmt;
		public IncreaseanddecreasestmtContext increaseanddecreasestmt;
		public IfstmtContext ifstmt;
		public WhilestmtContext whilestmt;
		public ForstmtContext forstmt;
		public SwitchstmtContext switchstmt;
		public GuardstmtContext guardstmt;
		public TransferstmtContext transferstmt;
		public DeclarevectorstmtContext declarevectorstmt;
		public AppendstmtContext appendstmt;
		public RemovelaststmtContext removelaststmt;
		public RemovestmtContext removestmt;
		public PrintstmtContext printstmt() {
			return getRuleContext(PrintstmtContext.class,0);
		}
		public DeclarestmtContext declarestmt() {
			return getRuleContext(DeclarestmtContext.class,0);
		}
		public ConstantstmtContext constantstmt() {
			return getRuleContext(ConstantstmtContext.class,0);
		}
		public AssignationstmtContext assignationstmt() {
			return getRuleContext(AssignationstmtContext.class,0);
		}
		public IncreaseanddecreasestmtContext increaseanddecreasestmt() {
			return getRuleContext(IncreaseanddecreasestmtContext.class,0);
		}
		public IfstmtContext ifstmt() {
			return getRuleContext(IfstmtContext.class,0);
		}
		public WhilestmtContext whilestmt() {
			return getRuleContext(WhilestmtContext.class,0);
		}
		public ForstmtContext forstmt() {
			return getRuleContext(ForstmtContext.class,0);
		}
		public SwitchstmtContext switchstmt() {
			return getRuleContext(SwitchstmtContext.class,0);
		}
		public GuardstmtContext guardstmt() {
			return getRuleContext(GuardstmtContext.class,0);
		}
		public TransferstmtContext transferstmt() {
			return getRuleContext(TransferstmtContext.class,0);
		}
		public DeclarevectorstmtContext declarevectorstmt() {
			return getRuleContext(DeclarevectorstmtContext.class,0);
		}
		public AppendstmtContext appendstmt() {
			return getRuleContext(AppendstmtContext.class,0);
		}
		public RemovelaststmtContext removelaststmt() {
			return getRuleContext(RemovelaststmtContext.class,0);
		}
		public RemovestmtContext removestmt() {
			return getRuleContext(RemovestmtContext.class,0);
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
			setState(110);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(65);
				((InstructionContext)_localctx).printstmt = printstmt();
				 _localctx.inst = ((InstructionContext)_localctx).printstmt.prnt
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(68);
				((InstructionContext)_localctx).declarestmt = declarestmt();
				_localctx.inst = ((InstructionContext)_localctx).declarestmt.dec
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(71);
				((InstructionContext)_localctx).constantstmt = constantstmt();
				_localctx.inst = ((InstructionContext)_localctx).constantstmt.const
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(74);
				((InstructionContext)_localctx).assignationstmt = assignationstmt();
				_localctx.inst = ((InstructionContext)_localctx).assignationstmt.assign
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(77);
				((InstructionContext)_localctx).increaseanddecreasestmt = increaseanddecreasestmt();
				_localctx.inst = ((InstructionContext)_localctx).increaseanddecreasestmt.increasedecrease
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(80);
				((InstructionContext)_localctx).ifstmt = ifstmt();
				 _localctx.inst = ((InstructionContext)_localctx).ifstmt.ift 
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(83);
				((InstructionContext)_localctx).whilestmt = whilestmt();
				_localctx.inst = ((InstructionContext)_localctx).whilestmt.while
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(86);
				((InstructionContext)_localctx).forstmt = forstmt();
				_localctx.inst = ((InstructionContext)_localctx).forstmt.for
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(89);
				((InstructionContext)_localctx).switchstmt = switchstmt();
				_localctx.inst = ((InstructionContext)_localctx).switchstmt.switch
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(92);
				((InstructionContext)_localctx).guardstmt = guardstmt();
				_localctx.inst = ((InstructionContext)_localctx).guardstmt.gua
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(95);
				((InstructionContext)_localctx).transferstmt = transferstmt();
				_localctx.inst = ((InstructionContext)_localctx).transferstmt.tran
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(98);
				((InstructionContext)_localctx).declarevectorstmt = declarevectorstmt();
				_localctx.inst = ((InstructionContext)_localctx).declarevectorstmt.decvec
				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(101);
				((InstructionContext)_localctx).appendstmt = appendstmt();
				_localctx.inst = ((InstructionContext)_localctx).appendstmt.app
				}
				break;
			case 14:
				enterOuterAlt(_localctx, 14);
				{
				setState(104);
				((InstructionContext)_localctx).removelaststmt = removelaststmt();
				_localctx.inst = ((InstructionContext)_localctx).removelaststmt.removl
				}
				break;
			case 15:
				enterOuterAlt(_localctx, 15);
				{
				setState(107);
				((InstructionContext)_localctx).removestmt = removestmt();
				_localctx.inst = ((InstructionContext)_localctx).removestmt.remov
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
			setState(112);
			((PrintstmtContext)_localctx).PRINT = match(PRINT);
			setState(113);
			match(PARIZQ);
			setState(114);
			((PrintstmtContext)_localctx).expr = expr(0);
			setState(115);
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
			setState(139);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(118);
				((DeclarestmtContext)_localctx).VAR = match(VAR);
				setState(119);
				((DeclarestmtContext)_localctx).ID = match(ID);
				setState(120);
				match(DOUBLEPTS);
				setState(121);
				((DeclarestmtContext)_localctx).type = type();
				setState(122);
				match(IG);
				setState(123);
				((DeclarestmtContext)_localctx).expr = expr(0);
				_localctx.dec = instructions.NewTodeclare((((DeclarestmtContext)_localctx).VAR!=null?((DeclarestmtContext)_localctx).VAR.getLine():0), (((DeclarestmtContext)_localctx).VAR!=null?((DeclarestmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarestmtContext)_localctx).ID!=null?((DeclarestmtContext)_localctx).ID.getText():null), ((DeclarestmtContext)_localctx).type.t, ((DeclarestmtContext)_localctx).expr.e, false)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(126);
				((DeclarestmtContext)_localctx).VAR = match(VAR);
				setState(127);
				((DeclarestmtContext)_localctx).ID = match(ID);
				setState(128);
				match(IG);
				setState(129);
				((DeclarestmtContext)_localctx).expr = expr(0);
				_localctx.dec = instructions.NewTodeclare((((DeclarestmtContext)_localctx).VAR!=null?((DeclarestmtContext)_localctx).VAR.getLine():0), (((DeclarestmtContext)_localctx).VAR!=null?((DeclarestmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarestmtContext)_localctx).ID!=null?((DeclarestmtContext)_localctx).ID.getText():null), environment.NULL, ((DeclarestmtContext)_localctx).expr.e, false)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(132);
				((DeclarestmtContext)_localctx).VAR = match(VAR);
				setState(133);
				((DeclarestmtContext)_localctx).ID = match(ID);
				setState(134);
				match(DOUBLEPTS);
				setState(135);
				((DeclarestmtContext)_localctx).type = type();
				setState(136);
				match(QUESTION);
				_localctx.dec = instructions.NewTodeclare((((DeclarestmtContext)_localctx).VAR!=null?((DeclarestmtContext)_localctx).VAR.getLine():0), (((DeclarestmtContext)_localctx).VAR!=null?((DeclarestmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarestmtContext)_localctx).ID!=null?((DeclarestmtContext)_localctx).ID.getText():null), ((DeclarestmtContext)_localctx).type.t, nil, false)
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
		public interfaces.Instruction const;
		public Token LET;
		public Token ID;
		public TypeContext type;
		public ExprContext expr;
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
			setState(155);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(141);
				((ConstantstmtContext)_localctx).LET = match(LET);
				setState(142);
				((ConstantstmtContext)_localctx).ID = match(ID);
				setState(143);
				match(DOUBLEPTS);
				setState(144);
				((ConstantstmtContext)_localctx).type = type();
				setState(145);
				match(IG);
				setState(146);
				((ConstantstmtContext)_localctx).expr = expr(0);
				_localctx.const = instructions.NewTodeclare((((ConstantstmtContext)_localctx).LET!=null?((ConstantstmtContext)_localctx).LET.getLine():0), (((ConstantstmtContext)_localctx).LET!=null?((ConstantstmtContext)_localctx).LET.getCharPositionInLine():0), (((ConstantstmtContext)_localctx).ID!=null?((ConstantstmtContext)_localctx).ID.getText():null), ((ConstantstmtContext)_localctx).type.t, ((ConstantstmtContext)_localctx).expr.e, true)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(149);
				((ConstantstmtContext)_localctx).LET = match(LET);
				setState(150);
				((ConstantstmtContext)_localctx).ID = match(ID);
				setState(151);
				match(IG);
				setState(152);
				((ConstantstmtContext)_localctx).expr = expr(0);
				_localctx.const = instructions.NewTodeclare((((ConstantstmtContext)_localctx).LET!=null?((ConstantstmtContext)_localctx).LET.getLine():0), (((ConstantstmtContext)_localctx).LET!=null?((ConstantstmtContext)_localctx).LET.getCharPositionInLine():0), (((ConstantstmtContext)_localctx).ID!=null?((ConstantstmtContext)_localctx).ID.getText():null), environment.NULL, ((ConstantstmtContext)_localctx).expr.e, true)
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

	public static class BlockelifsContext extends ParserRuleContext {
		public []interface{} blkef;
		public IfstmtContext ifstmt;
		public List<IfstmtContext> elif = new ArrayList<IfstmtContext>();
		public List<IfstmtContext> ifstmt() {
			return getRuleContexts(IfstmtContext.class);
		}
		public IfstmtContext ifstmt(int i) {
			return getRuleContext(IfstmtContext.class,i);
		}
		public BlockelifsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_blockelifs; }
	}

	public final BlockelifsContext blockelifs() throws RecognitionException {
		BlockelifsContext _localctx = new BlockelifsContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_blockelifs);

		_localctx.blkef = []interface{}{}
		    var listifs []IIfstmtContext
		    
		  
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(158); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(157);
					((BlockelifsContext)_localctx).ifstmt = ifstmt();
					((BlockelifsContext)_localctx).elif.add(((BlockelifsContext)_localctx).ifstmt);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(160); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
			   
			        listifs = localctx.(*BlockelifsContext).GetElif()
			        for _, a := range listifs {
			            _localctx.blkef = append(_localctx.blkef, a.GetIft())
			            
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

	public static class IfstmtContext extends ParserRuleContext {
		public interfaces.Instruction ift;
		public []interface{} el;
		public interfaces.Instruction else;
		public Token IF;
		public ExprContext expr;
		public BlockContext ifb;
		public BlockContext ifelseblck;
		public BlockContext elseifblck;
		public BlockContext elif;
		public BlockelifsContext blockelifs;
		public TerminalNode IF() { return getToken(GrammarParser.IF, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public List<TerminalNode> LLAVEIZQ() { return getTokens(GrammarParser.LLAVEIZQ); }
		public TerminalNode LLAVEIZQ(int i) {
			return getToken(GrammarParser.LLAVEIZQ, i);
		}
		public List<TerminalNode> LLAVEDER() { return getTokens(GrammarParser.LLAVEDER); }
		public TerminalNode LLAVEDER(int i) {
			return getToken(GrammarParser.LLAVEDER, i);
		}
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public TerminalNode ELSE() { return getToken(GrammarParser.ELSE, 0); }
		public BlockelifsContext blockelifs() {
			return getRuleContext(BlockelifsContext.class,0);
		}
		public IfstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifstmt; }
	}

	public final IfstmtContext ifstmt() throws RecognitionException {
		IfstmtContext _localctx = new IfstmtContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_ifstmt);
		try {
			setState(191);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(164);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(165);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(166);
				match(LLAVEIZQ);
				setState(167);
				((IfstmtContext)_localctx).ifb = block();
				setState(168);
				match(LLAVEDER);
				 ((IfstmtContext)_localctx).ift =  instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).ifb.blk, nil); 
				                                        _localctx.el = ((IfstmtContext)_localctx).ifb.blk
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(171);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(172);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(173);
				match(LLAVEIZQ);
				setState(174);
				((IfstmtContext)_localctx).ifelseblck = block();
				setState(175);
				match(LLAVEDER);
				setState(176);
				match(ELSE);
				setState(177);
				match(LLAVEIZQ);
				setState(178);
				((IfstmtContext)_localctx).elseifblck = block();
				setState(179);
				match(LLAVEDER);
				 ((IfstmtContext)_localctx).ift =  instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).ifelseblck.blk, ((IfstmtContext)_localctx).elseifblck.blk); 
				                                                                                        ((IfstmtContext)_localctx).el =  ((IfstmtContext)_localctx).ifelseblck.blk;
				                                                                                        
				                                                                                        
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(182);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(183);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(184);
				match(LLAVEIZQ);
				setState(185);
				((IfstmtContext)_localctx).elif = block();
				setState(186);
				match(LLAVEDER);
				setState(187);
				match(ELSE);
				setState(188);
				((IfstmtContext)_localctx).blockelifs = blockelifs();
				 ((IfstmtContext)_localctx).ift =  instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).elif.blk, ((IfstmtContext)_localctx).blockelifs.blkef); 
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

	public static class SwitchstmtContext extends ParserRuleContext {
		public interfaces.Instruction switch;
		public Token SWITCH;
		public ExprContext expr;
		public BlockContext block;
		public CasestmtContext casestmt;
		public TerminalNode SWITCH() { return getToken(GrammarParser.SWITCH, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(GrammarParser.LLAVEIZQ, 0); }
		public TerminalNode DEFAULT() { return getToken(GrammarParser.DEFAULT, 0); }
		public TerminalNode DOUBLEPTS() { return getToken(GrammarParser.DOUBLEPTS, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(GrammarParser.LLAVEDER, 0); }
		public CasestmtContext casestmt() {
			return getRuleContext(CasestmtContext.class,0);
		}
		public SwitchstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switchstmt; }
	}

	public final SwitchstmtContext switchstmt() throws RecognitionException {
		SwitchstmtContext _localctx = new SwitchstmtContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_switchstmt);
		try {
			setState(209);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(193);
				((SwitchstmtContext)_localctx).SWITCH = match(SWITCH);
				setState(194);
				((SwitchstmtContext)_localctx).expr = expr(0);
				setState(195);
				match(LLAVEIZQ);
				setState(196);
				match(DEFAULT);
				setState(197);
				match(DOUBLEPTS);
				setState(198);
				((SwitchstmtContext)_localctx).block = block();
				setState(199);
				match(LLAVEDER);
				 ((SwitchstmtContext)_localctx).switch =  instructions.NewSwitch((((SwitchstmtContext)_localctx).SWITCH!=null?((SwitchstmtContext)_localctx).SWITCH.getLine():0), (((SwitchstmtContext)_localctx).SWITCH!=null?((SwitchstmtContext)_localctx).SWITCH.getCharPositionInLine():0), ((SwitchstmtContext)_localctx).expr.e, nil,  nil, ((SwitchstmtContext)_localctx).block.blk);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(202);
				((SwitchstmtContext)_localctx).SWITCH = match(SWITCH);
				setState(203);
				((SwitchstmtContext)_localctx).expr = expr(0);
				setState(204);
				match(LLAVEIZQ);
				setState(205);
				((SwitchstmtContext)_localctx).casestmt = casestmt();
				setState(206);
				match(LLAVEDER);
				((SwitchstmtContext)_localctx).switch =  instructions.NewSwitch((((SwitchstmtContext)_localctx).SWITCH!=null?((SwitchstmtContext)_localctx).SWITCH.getLine():0), (((SwitchstmtContext)_localctx).SWITCH!=null?((SwitchstmtContext)_localctx).SWITCH.getCharPositionInLine():0), ((SwitchstmtContext)_localctx).expr.e, nil, ((SwitchstmtContext)_localctx).casestmt.cas, nil);
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

	public static class CasestmtContext extends ParserRuleContext {
		public interfaces.Instruction cas;
		public Token CASE;
		public ExprContext expr;
		public BlockContext block;
		public BlockcasesContext blockcases;
		public BlockContext sen;
		public BlockContext def;
		public TerminalNode CASE() { return getToken(GrammarParser.CASE, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public List<TerminalNode> DOUBLEPTS() { return getTokens(GrammarParser.DOUBLEPTS); }
		public TerminalNode DOUBLEPTS(int i) {
			return getToken(GrammarParser.DOUBLEPTS, i);
		}
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public BlockcasesContext blockcases() {
			return getRuleContext(BlockcasesContext.class,0);
		}
		public TerminalNode DEFAULT() { return getToken(GrammarParser.DEFAULT, 0); }
		public CasestmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_casestmt; }
	}

	public final CasestmtContext casestmt() throws RecognitionException {
		CasestmtContext _localctx = new CasestmtContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_casestmt);
		try {
			setState(227);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(211);
				((CasestmtContext)_localctx).CASE = match(CASE);
				setState(212);
				((CasestmtContext)_localctx).expr = expr(0);
				setState(213);
				match(DOUBLEPTS);
				setState(214);
				((CasestmtContext)_localctx).block = block();
				setState(215);
				((CasestmtContext)_localctx).blockcases = blockcases();
				 _localctx.cas = instructions.NewCase((((CasestmtContext)_localctx).CASE!=null?((CasestmtContext)_localctx).CASE.getLine():0),(((CasestmtContext)_localctx).CASE!=null?((CasestmtContext)_localctx).CASE.getCharPositionInLine():0), ((CasestmtContext)_localctx).expr.e,((CasestmtContext)_localctx).block.blk, ((CasestmtContext)_localctx).blockcases.blkcase)
				                                fmt.Println("entro")
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(218);
				((CasestmtContext)_localctx).CASE = match(CASE);
				setState(219);
				((CasestmtContext)_localctx).expr = expr(0);
				setState(220);
				match(DOUBLEPTS);
				setState(221);
				((CasestmtContext)_localctx).sen = block();
				setState(222);
				match(DEFAULT);
				setState(223);
				match(DOUBLEPTS);
				setState(224);
				((CasestmtContext)_localctx).def = block();
				 _localctx.cas = instructions.NewCase((((CasestmtContext)_localctx).CASE!=null?((CasestmtContext)_localctx).CASE.getLine():0),(((CasestmtContext)_localctx).CASE!=null?((CasestmtContext)_localctx).CASE.getCharPositionInLine():0), ((CasestmtContext)_localctx).expr.e,((CasestmtContext)_localctx).sen.blk, ((CasestmtContext)_localctx).def.blk)
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

	public static class BlockcasesContext extends ParserRuleContext {
		public []interface{} blkcase;
		public CasestmtContext casestmt;
		public List<CasestmtContext> casedef = new ArrayList<CasestmtContext>();
		public List<CasestmtContext> casestmt() {
			return getRuleContexts(CasestmtContext.class);
		}
		public CasestmtContext casestmt(int i) {
			return getRuleContext(CasestmtContext.class,i);
		}
		public BlockcasesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_blockcases; }
	}

	public final BlockcasesContext blockcases() throws RecognitionException {
		BlockcasesContext _localctx = new BlockcasesContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_blockcases);

		    _localctx.blkcase = []interface{}{}
		    var listcases []ICasestmtContext
		  
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(230); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(229);
					((BlockcasesContext)_localctx).casestmt = casestmt();
					((BlockcasesContext)_localctx).casedef.add(((BlockcasesContext)_localctx).casestmt);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(232); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );

			        listcases = localctx.(*BlockcasesContext).GetCasedef()
			        for _, a := range listcases {
			            _localctx.blkcase = append(_localctx.blkcase, a.GetCas())
			      
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

	public static class AssignationstmtContext extends ParserRuleContext {
		public interfaces.Instruction assign;
		public Token ID;
		public Token IG;
		public ExprContext expr;
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode IG() { return getToken(GrammarParser.IG, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public AssignationstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignationstmt; }
	}

	public final AssignationstmtContext assignationstmt() throws RecognitionException {
		AssignationstmtContext _localctx = new AssignationstmtContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_assignationstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(236);
			((AssignationstmtContext)_localctx).ID = match(ID);
			setState(237);
			((AssignationstmtContext)_localctx).IG = match(IG);
			setState(238);
			((AssignationstmtContext)_localctx).expr = expr(0);
			_localctx.assign = instructions.NewAssignation((((AssignationstmtContext)_localctx).IG!=null?((AssignationstmtContext)_localctx).IG.getLine():0), (((AssignationstmtContext)_localctx).IG!=null?((AssignationstmtContext)_localctx).IG.getCharPositionInLine():0),  (((AssignationstmtContext)_localctx).ID!=null?((AssignationstmtContext)_localctx).ID.getText():null), ((AssignationstmtContext)_localctx).expr.e)
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
		public Token WHILE;
		public ExprContext expr;
		public BlockContext block;
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
		enterRule(_localctx, 24, RULE_whilestmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(241);
			((WhilestmtContext)_localctx).WHILE = match(WHILE);
			setState(242);
			((WhilestmtContext)_localctx).expr = expr(0);
			setState(243);
			match(LLAVEIZQ);
			setState(244);
			((WhilestmtContext)_localctx).block = block();
			setState(245);
			match(LLAVEDER);
			_localctx.while = instructions.NewWhile((((WhilestmtContext)_localctx).WHILE!=null?((WhilestmtContext)_localctx).WHILE.getLine():0), (((WhilestmtContext)_localctx).WHILE!=null?((WhilestmtContext)_localctx).WHILE.getCharPositionInLine():0),  ((WhilestmtContext)_localctx).expr.e, ((WhilestmtContext)_localctx).block.blk)
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
		public Token FOR;
		public Token ID;
		public ExprContext expr;
		public BlockContext block;
		public ExprContext first;
		public ExprContext second;
		public TerminalNode FOR() { return getToken(GrammarParser.FOR, 0); }
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode IN() { return getToken(GrammarParser.IN, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode LLAVEIZQ() { return getToken(GrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(GrammarParser.LLAVEDER, 0); }
		public List<TerminalNode> POINT() { return getTokens(GrammarParser.POINT); }
		public TerminalNode POINT(int i) {
			return getToken(GrammarParser.POINT, i);
		}
		public ForstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_forstmt; }
	}

	public final ForstmtContext forstmt() throws RecognitionException {
		ForstmtContext _localctx = new ForstmtContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_forstmt);
		try {
			setState(269);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(248);
				((ForstmtContext)_localctx).FOR = match(FOR);
				setState(249);
				((ForstmtContext)_localctx).ID = match(ID);
				setState(250);
				match(IN);
				setState(251);
				((ForstmtContext)_localctx).expr = expr(0);
				setState(252);
				match(LLAVEIZQ);
				setState(253);
				((ForstmtContext)_localctx).block = block();
				setState(254);
				match(LLAVEDER);
				_localctx.for = instructions.NewFor((((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getLine():0), (((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getCharPositionInLine():0), (((ForstmtContext)_localctx).ID!=null?((ForstmtContext)_localctx).ID.getText():null), ((ForstmtContext)_localctx).expr.e, nil, ((ForstmtContext)_localctx).block.blk)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(257);
				((ForstmtContext)_localctx).FOR = match(FOR);
				setState(258);
				((ForstmtContext)_localctx).ID = match(ID);
				setState(259);
				match(IN);
				setState(260);
				((ForstmtContext)_localctx).first = expr(0);
				setState(261);
				match(POINT);
				setState(262);
				match(POINT);
				setState(263);
				((ForstmtContext)_localctx).second = expr(0);
				setState(264);
				match(LLAVEIZQ);
				setState(265);
				((ForstmtContext)_localctx).block = block();
				setState(266);
				match(LLAVEDER);
				_localctx.for = instructions.NewFor((((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getLine():0), (((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getCharPositionInLine():0), (((ForstmtContext)_localctx).ID!=null?((ForstmtContext)_localctx).ID.getText():null), ((ForstmtContext)_localctx).first.e, ((ForstmtContext)_localctx).second.e, ((ForstmtContext)_localctx).block.blk)
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

	public static class GuardstmtContext extends ParserRuleContext {
		public interfaces.Instruction gua;
		public Token GUARD;
		public ExprContext expr;
		public BlockContext block;
		public Token r;
		public TerminalNode GUARD() { return getToken(GrammarParser.GUARD, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode ELSE() { return getToken(GrammarParser.ELSE, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(GrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(GrammarParser.LLAVEDER, 0); }
		public TerminalNode CONTINUE() { return getToken(GrammarParser.CONTINUE, 0); }
		public TerminalNode RETURN() { return getToken(GrammarParser.RETURN, 0); }
		public TerminalNode BREAK() { return getToken(GrammarParser.BREAK, 0); }
		public GuardstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_guardstmt; }
	}

	public final GuardstmtContext guardstmt() throws RecognitionException {
		GuardstmtContext _localctx = new GuardstmtContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_guardstmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(271);
			((GuardstmtContext)_localctx).GUARD = match(GUARD);
			setState(272);
			((GuardstmtContext)_localctx).expr = expr(0);
			setState(273);
			match(ELSE);
			setState(274);
			match(LLAVEIZQ);
			setState(275);
			((GuardstmtContext)_localctx).block = block();
			setState(276);
			((GuardstmtContext)_localctx).r = _input.LT(1);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << CONTINUE) | (1L << RETURN) | (1L << BREAK))) != 0)) ) {
				((GuardstmtContext)_localctx).r = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(277);
			match(LLAVEDER);
			_localctx.gua = instructions.NewGuard((((GuardstmtContext)_localctx).GUARD!=null?((GuardstmtContext)_localctx).GUARD.getLine():0), (((GuardstmtContext)_localctx).GUARD!=null?((GuardstmtContext)_localctx).GUARD.getCharPositionInLine():0), ((GuardstmtContext)_localctx).expr.e, ((GuardstmtContext)_localctx).block.blk, (((GuardstmtContext)_localctx).r!=null?((GuardstmtContext)_localctx).r.getText():null))
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

	public static class TransferstmtContext extends ParserRuleContext {
		public interfaces.Instruction tran;
		public Token RETURN;
		public ExprContext expr;
		public Token CONTINUE;
		public Token BREAK;
		public TerminalNode RETURN() { return getToken(GrammarParser.RETURN, 0); }
		public TerminalNode PTCOMA() { return getToken(GrammarParser.PTCOMA, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode CONTINUE() { return getToken(GrammarParser.CONTINUE, 0); }
		public TerminalNode BREAK() { return getToken(GrammarParser.BREAK, 0); }
		public TransferstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_transferstmt; }
	}

	public final TransferstmtContext transferstmt() throws RecognitionException {
		TransferstmtContext _localctx = new TransferstmtContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_transferstmt);
		try {
			setState(294);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(280);
				((TransferstmtContext)_localctx).RETURN = match(RETURN);
				setState(281);
				match(PTCOMA);
				_localctx.tran = instructions.NewReturnIn((((TransferstmtContext)_localctx).RETURN!=null?((TransferstmtContext)_localctx).RETURN.getLine():0), (((TransferstmtContext)_localctx).RETURN!=null?((TransferstmtContext)_localctx).RETURN.getCharPositionInLine():0), nil)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(283);
				((TransferstmtContext)_localctx).RETURN = match(RETURN);
				setState(284);
				((TransferstmtContext)_localctx).expr = expr(0);
				setState(285);
				match(PTCOMA);
				_localctx.tran = instructions.NewReturnIn((((TransferstmtContext)_localctx).RETURN!=null?((TransferstmtContext)_localctx).RETURN.getLine():0), (((TransferstmtContext)_localctx).RETURN!=null?((TransferstmtContext)_localctx).RETURN.getCharPositionInLine():0), ((TransferstmtContext)_localctx).expr.e)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(288);
				((TransferstmtContext)_localctx).CONTINUE = match(CONTINUE);
				setState(289);
				match(PTCOMA);
				_localctx.tran = instructions.NewContinue((((TransferstmtContext)_localctx).CONTINUE!=null?((TransferstmtContext)_localctx).CONTINUE.getLine():0), (((TransferstmtContext)_localctx).CONTINUE!=null?((TransferstmtContext)_localctx).CONTINUE.getCharPositionInLine():0))
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(291);
				((TransferstmtContext)_localctx).BREAK = match(BREAK);
				setState(292);
				match(PTCOMA);
				_localctx.tran = instructions.NewBreak((((TransferstmtContext)_localctx).BREAK!=null?((TransferstmtContext)_localctx).BREAK.getLine():0), (((TransferstmtContext)_localctx).BREAK!=null?((TransferstmtContext)_localctx).BREAK.getCharPositionInLine():0))
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

	public static class DeclarevectorstmtContext extends ParserRuleContext {
		public interfaces.Instruction decvec;
		public Token VAR;
		public Token ID;
		public TypeContext type;
		public ListParamsContext listParams;
		public TerminalNode VAR() { return getToken(GrammarParser.VAR, 0); }
		public List<TerminalNode> ID() { return getTokens(GrammarParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(GrammarParser.ID, i);
		}
		public TerminalNode DOUBLEPTS() { return getToken(GrammarParser.DOUBLEPTS, 0); }
		public List<TerminalNode> CORCHETEIZQ() { return getTokens(GrammarParser.CORCHETEIZQ); }
		public TerminalNode CORCHETEIZQ(int i) {
			return getToken(GrammarParser.CORCHETEIZQ, i);
		}
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public List<TerminalNode> CORCHETEDER() { return getTokens(GrammarParser.CORCHETEDER); }
		public TerminalNode CORCHETEDER(int i) {
			return getToken(GrammarParser.CORCHETEDER, i);
		}
		public TerminalNode IG() { return getToken(GrammarParser.IG, 0); }
		public ListParamsContext listParams() {
			return getRuleContext(ListParamsContext.class,0);
		}
		public TerminalNode PARIZQ() { return getToken(GrammarParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(GrammarParser.PARDER, 0); }
		public DeclarevectorstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declarevectorstmt; }
	}

	public final DeclarevectorstmtContext declarevectorstmt() throws RecognitionException {
		DeclarevectorstmtContext _localctx = new DeclarevectorstmtContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_declarevectorstmt);
		try {
			setState(337);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(296);
				((DeclarevectorstmtContext)_localctx).VAR = match(VAR);
				setState(297);
				((DeclarevectorstmtContext)_localctx).ID = match(ID);
				setState(298);
				match(DOUBLEPTS);
				setState(299);
				match(CORCHETEIZQ);
				setState(300);
				((DeclarevectorstmtContext)_localctx).type = type();
				setState(301);
				match(CORCHETEDER);
				setState(302);
				match(IG);
				setState(303);
				match(CORCHETEIZQ);
				setState(304);
				((DeclarevectorstmtContext)_localctx).listParams = listParams(0);
				setState(305);
				match(CORCHETEDER);
				_localctx.decvec = instructions.NewToDecalreVector((((DeclarevectorstmtContext)_localctx).VAR!=null?((DeclarevectorstmtContext)_localctx).VAR.getLine():0), (((DeclarevectorstmtContext)_localctx).VAR!=null?((DeclarevectorstmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarevectorstmtContext)_localctx).ID!=null?((DeclarevectorstmtContext)_localctx).ID.getText():null), ((DeclarevectorstmtContext)_localctx).type.t, ((DeclarevectorstmtContext)_localctx).listParams.l)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(308);
				((DeclarevectorstmtContext)_localctx).VAR = match(VAR);
				setState(309);
				((DeclarevectorstmtContext)_localctx).ID = match(ID);
				setState(310);
				match(DOUBLEPTS);
				setState(311);
				match(CORCHETEIZQ);
				setState(312);
				((DeclarevectorstmtContext)_localctx).type = type();
				setState(313);
				match(CORCHETEDER);
				setState(314);
				match(IG);
				setState(315);
				match(CORCHETEIZQ);
				setState(316);
				match(CORCHETEDER);
				_localctx.decvec = instructions.NewToDecalreVector((((DeclarevectorstmtContext)_localctx).VAR!=null?((DeclarevectorstmtContext)_localctx).VAR.getLine():0), (((DeclarevectorstmtContext)_localctx).VAR!=null?((DeclarevectorstmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarevectorstmtContext)_localctx).ID!=null?((DeclarevectorstmtContext)_localctx).ID.getText():null), ((DeclarevectorstmtContext)_localctx).type.t, nil)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(319);
				match(VAR);
				setState(320);
				match(ID);
				setState(321);
				match(IG);
				setState(322);
				match(CORCHETEIZQ);
				setState(323);
				type();
				setState(324);
				match(CORCHETEDER);
				setState(325);
				match(PARIZQ);
				setState(326);
				match(PARDER);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(328);
				match(VAR);
				setState(329);
				match(ID);
				setState(330);
				match(DOUBLEPTS);
				setState(331);
				match(CORCHETEIZQ);
				setState(332);
				type();
				setState(333);
				match(CORCHETEDER);
				setState(334);
				match(IG);
				setState(335);
				match(ID);
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

	public static class AppendstmtContext extends ParserRuleContext {
		public interfaces.Instruction app;
		public Token ID;
		public ExprContext expr;
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode POINT() { return getToken(GrammarParser.POINT, 0); }
		public TerminalNode APPEND() { return getToken(GrammarParser.APPEND, 0); }
		public TerminalNode PARIZQ() { return getToken(GrammarParser.PARIZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(GrammarParser.PARDER, 0); }
		public AppendstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_appendstmt; }
	}

	public final AppendstmtContext appendstmt() throws RecognitionException {
		AppendstmtContext _localctx = new AppendstmtContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_appendstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(339);
			((AppendstmtContext)_localctx).ID = match(ID);
			setState(340);
			match(POINT);
			setState(341);
			match(APPEND);
			setState(342);
			match(PARIZQ);
			setState(343);
			((AppendstmtContext)_localctx).expr = expr(0);
			setState(344);
			match(PARDER);
			_localctx.app = instructions.NewAppend((((AppendstmtContext)_localctx).ID!=null?((AppendstmtContext)_localctx).ID.getLine():0), (((AppendstmtContext)_localctx).ID!=null?((AppendstmtContext)_localctx).ID.getCharPositionInLine():0), (((AppendstmtContext)_localctx).ID!=null?((AppendstmtContext)_localctx).ID.getText():null), ((AppendstmtContext)_localctx).expr.e)
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

	public static class RemovelaststmtContext extends ParserRuleContext {
		public interfaces.Instruction removl;
		public Token ID;
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode POINT() { return getToken(GrammarParser.POINT, 0); }
		public TerminalNode REMOVELAST() { return getToken(GrammarParser.REMOVELAST, 0); }
		public TerminalNode PARIZQ() { return getToken(GrammarParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(GrammarParser.PARDER, 0); }
		public RemovelaststmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_removelaststmt; }
	}

	public final RemovelaststmtContext removelaststmt() throws RecognitionException {
		RemovelaststmtContext _localctx = new RemovelaststmtContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_removelaststmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(347);
			((RemovelaststmtContext)_localctx).ID = match(ID);
			setState(348);
			match(POINT);
			setState(349);
			match(REMOVELAST);
			setState(350);
			match(PARIZQ);
			setState(351);
			match(PARDER);
			_localctx.removl = instructions.NewRemoveLast((((RemovelaststmtContext)_localctx).ID!=null?((RemovelaststmtContext)_localctx).ID.getLine():0), (((RemovelaststmtContext)_localctx).ID!=null?((RemovelaststmtContext)_localctx).ID.getCharPositionInLine():0), (((RemovelaststmtContext)_localctx).ID!=null?((RemovelaststmtContext)_localctx).ID.getText():null))
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

	public static class RemovestmtContext extends ParserRuleContext {
		public interfaces.Instruction remov;
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode POINT() { return getToken(GrammarParser.POINT, 0); }
		public TerminalNode REMOVE() { return getToken(GrammarParser.REMOVE, 0); }
		public TerminalNode PARIZQ() { return getToken(GrammarParser.PARIZQ, 0); }
		public TerminalNode AT() { return getToken(GrammarParser.AT, 0); }
		public TerminalNode DOUBLEPTS() { return getToken(GrammarParser.DOUBLEPTS, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(GrammarParser.PARDER, 0); }
		public RemovestmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_removestmt; }
	}

	public final RemovestmtContext removestmt() throws RecognitionException {
		RemovestmtContext _localctx = new RemovestmtContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_removestmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(354);
			match(ID);
			setState(355);
			match(POINT);
			setState(356);
			match(REMOVE);
			setState(357);
			match(PARIZQ);
			setState(358);
			match(AT);
			setState(359);
			match(DOUBLEPTS);
			setState(360);
			expr(0);
			setState(361);
			match(PARDER);
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

	public static class EmptvecstmtContext extends ParserRuleContext {
		public interfaces.Expression emptyvec;
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode POINT() { return getToken(GrammarParser.POINT, 0); }
		public TerminalNode ISEMPTY() { return getToken(GrammarParser.ISEMPTY, 0); }
		public EmptvecstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_emptvecstmt; }
	}

	public final EmptvecstmtContext emptvecstmt() throws RecognitionException {
		EmptvecstmtContext _localctx = new EmptvecstmtContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_emptvecstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(363);
			match(ID);
			setState(364);
			match(POINT);
			setState(365);
			match(ISEMPTY);
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

	public static class CountvecstmtContext extends ParserRuleContext {
		public interfaces.Expression count;
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode POINT() { return getToken(GrammarParser.POINT, 0); }
		public TerminalNode COUNT() { return getToken(GrammarParser.COUNT, 0); }
		public CountvecstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_countvecstmt; }
	}

	public final CountvecstmtContext countvecstmt() throws RecognitionException {
		CountvecstmtContext _localctx = new CountvecstmtContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_countvecstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(367);
			match(ID);
			setState(368);
			match(POINT);
			setState(369);
			match(COUNT);
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

	public static class ListParamsContext extends ParserRuleContext {
		public []interface{} l;
		public ListParamsContext list;
		public ExprContext expr;
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode COMA() { return getToken(GrammarParser.COMA, 0); }
		public ListParamsContext listParams() {
			return getRuleContext(ListParamsContext.class,0);
		}
		public ListParamsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listParams; }
	}

	public final ListParamsContext listParams() throws RecognitionException {
		return listParams(0);
	}

	private ListParamsContext listParams(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ListParamsContext _localctx = new ListParamsContext(_ctx, _parentState);
		ListParamsContext _prevctx = _localctx;
		int _startState = 44;
		enterRecursionRule(_localctx, 44, RULE_listParams, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(372);
			((ListParamsContext)_localctx).expr = expr(0);

			            _localctx.l = []interface{}{}
			            _localctx.l = append(_localctx.l, ((ListParamsContext)_localctx).expr.e)
			        
			}
			_ctx.stop = _input.LT(-1);
			setState(382);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,12,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ListParamsContext(_parentctx, _parentState);
					_localctx.list = _prevctx;
					_localctx.list = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_listParams);
					setState(375);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(376);
					match(COMA);
					setState(377);
					((ListParamsContext)_localctx).expr = expr(0);

					                                          var arr []interface{}
					                                          arr = append(((ListParamsContext)_localctx).list.l, ((ListParamsContext)_localctx).expr.e)
					                                          _localctx.l = arr
					                                      
					}
					} 
				}
				setState(384);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,12,_ctx);
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

	public static class ExprContext extends ParserRuleContext {
		public interfaces.Expression e;
		public ExprContext left;
		public Token op;
		public ExprContext expr;
		public Token NUMBER;
		public Token STRING;
		public Token TRUE;
		public Token FALSE;
		public Token NIL;
		public AccessstmtContext accessstmt;
		public ExprContext right;
		public TerminalNode NOT() { return getToken(GrammarParser.NOT, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode SUB() { return getToken(GrammarParser.SUB, 0); }
		public TerminalNode PARIZQ() { return getToken(GrammarParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(GrammarParser.PARDER, 0); }
		public TerminalNode NUMBER() { return getToken(GrammarParser.NUMBER, 0); }
		public TerminalNode STRING() { return getToken(GrammarParser.STRING, 0); }
		public TerminalNode TRUE() { return getToken(GrammarParser.TRUE, 0); }
		public TerminalNode FALSE() { return getToken(GrammarParser.FALSE, 0); }
		public TerminalNode NIL() { return getToken(GrammarParser.NIL, 0); }
		public AccessstmtContext accessstmt() {
			return getRuleContext(AccessstmtContext.class,0);
		}
		public TerminalNode MUL() { return getToken(GrammarParser.MUL, 0); }
		public TerminalNode DIV() { return getToken(GrammarParser.DIV, 0); }
		public TerminalNode ADD() { return getToken(GrammarParser.ADD, 0); }
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
		int _startState = 46;
		enterRecursionRule(_localctx, 46, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(412);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NOT:
				{
				setState(386);
				((ExprContext)_localctx).op = match(NOT);
				setState(387);
				((ExprContext)_localctx).left = ((ExprContext)_localctx).expr = expr(9);
				 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).left.e) 
				}
				break;
			case SUB:
				{
				setState(390);
				((ExprContext)_localctx).op = match(SUB);
				setState(391);
				((ExprContext)_localctx).left = ((ExprContext)_localctx).expr = expr(8);
				 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e,  (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), nil)
				}
				break;
			case PARIZQ:
				{
				setState(394);
				match(PARIZQ);
				setState(395);
				((ExprContext)_localctx).expr = expr(0);
				setState(396);
				match(PARDER);
				 _localctx.e = ((ExprContext)_localctx).expr.e 
				}
				break;
			case NUMBER:
				{
				setState(399);
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
				setState(401);
				((ExprContext)_localctx).STRING = match(STRING);

				        str := (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getText():null)
				        _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getLine():0), (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getCharPositionInLine():0), str[1:len(str)-1],environment.STRING)
				    
				}
				break;
			case TRUE:
				{
				setState(403);
				((ExprContext)_localctx).TRUE = match(TRUE);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).TRUE!=null?((ExprContext)_localctx).TRUE.getLine():0), (((ExprContext)_localctx).TRUE!=null?((ExprContext)_localctx).TRUE.getCharPositionInLine():0), true, environment.BOOLEAN) 
				}
				break;
			case FALSE:
				{
				setState(405);
				((ExprContext)_localctx).FALSE = match(FALSE);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).FALSE!=null?((ExprContext)_localctx).FALSE.getLine():0), (((ExprContext)_localctx).FALSE!=null?((ExprContext)_localctx).FALSE.getCharPositionInLine():0), false, environment.BOOLEAN) 
				}
				break;
			case NIL:
				{
				setState(407);
				((ExprContext)_localctx).NIL = match(NIL);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).NIL!=null?((ExprContext)_localctx).NIL.getLine():0), (((ExprContext)_localctx).NIL!=null?((ExprContext)_localctx).NIL.getCharPositionInLine():0), nil, environment.NULL) 
				}
				break;
			case ID:
				{
				setState(409);
				((ExprContext)_localctx).accessstmt = accessstmt();
				_localctx.e = ((ExprContext)_localctx).accessstmt.access
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(451);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(449);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
					case 1:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(414);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(415);
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
						setState(416);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(17);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 2:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(419);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(420);
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
						setState(421);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(16);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 3:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(424);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(425);
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
						setState(426);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(15);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 4:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(429);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(430);
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
						setState(431);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(14);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 5:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(434);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(435);
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
						setState(436);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(13);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 6:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(439);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(440);
						((ExprContext)_localctx).op = match(AND);
						setState(441);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(12);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 7:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(444);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(445);
						((ExprContext)_localctx).op = match(OR);
						setState(446);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(11);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					}
					} 
				}
				setState(453);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
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
		enterRule(_localctx, 48, RULE_accessstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(454);
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

	public static class IncreaseanddecreasestmtContext extends ParserRuleContext {
		public interfaces.Instruction increasedecrease;
		public Token ID;
		public Token IG_ADD;
		public ExprContext expr;
		public Token IG_SUB;
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode IG_ADD() { return getToken(GrammarParser.IG_ADD, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode IG_SUB() { return getToken(GrammarParser.IG_SUB, 0); }
		public IncreaseanddecreasestmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_increaseanddecreasestmt; }
	}

	public final IncreaseanddecreasestmtContext increaseanddecreasestmt() throws RecognitionException {
		IncreaseanddecreasestmtContext _localctx = new IncreaseanddecreasestmtContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_increaseanddecreasestmt);
		try {
			setState(467);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(457);
				((IncreaseanddecreasestmtContext)_localctx).ID = match(ID);
				setState(458);
				((IncreaseanddecreasestmtContext)_localctx).IG_ADD = match(IG_ADD);
				setState(459);
				((IncreaseanddecreasestmtContext)_localctx).expr = expr(0);
				_localctx.increasedecrease = instructions.NewIncreaseDecrease((((IncreaseanddecreasestmtContext)_localctx).ID!=null?((IncreaseanddecreasestmtContext)_localctx).ID.getLine():0), (((IncreaseanddecreasestmtContext)_localctx).ID!=null?((IncreaseanddecreasestmtContext)_localctx).ID.getCharPositionInLine():0), (((IncreaseanddecreasestmtContext)_localctx).ID!=null?((IncreaseanddecreasestmtContext)_localctx).ID.getText():null), (((IncreaseanddecreasestmtContext)_localctx).IG_ADD!=null?((IncreaseanddecreasestmtContext)_localctx).IG_ADD.getText():null), ((IncreaseanddecreasestmtContext)_localctx).expr.e)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(462);
				((IncreaseanddecreasestmtContext)_localctx).ID = match(ID);
				setState(463);
				((IncreaseanddecreasestmtContext)_localctx).IG_SUB = match(IG_SUB);
				setState(464);
				((IncreaseanddecreasestmtContext)_localctx).expr = expr(0);
				_localctx.increasedecrease = instructions.NewIncreaseDecrease((((IncreaseanddecreasestmtContext)_localctx).ID!=null?((IncreaseanddecreasestmtContext)_localctx).ID.getLine():0), (((IncreaseanddecreasestmtContext)_localctx).ID!=null?((IncreaseanddecreasestmtContext)_localctx).ID.getCharPositionInLine():0), (((IncreaseanddecreasestmtContext)_localctx).ID!=null?((IncreaseanddecreasestmtContext)_localctx).ID.getText():null), (((IncreaseanddecreasestmtContext)_localctx).IG_SUB!=null?((IncreaseanddecreasestmtContext)_localctx).IG_SUB.getText():null), ((IncreaseanddecreasestmtContext)_localctx).expr.e)
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
		enterRule(_localctx, 52, RULE_type);
		try {
			setState(479);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRINGS:
				enterOuterAlt(_localctx, 1);
				{
				setState(469);
				match(STRINGS);
				_localctx.t = environment.STRING
				}
				break;
			case INTS:
				enterOuterAlt(_localctx, 2);
				{
				setState(471);
				match(INTS);
				_localctx.t = environment.INTEGER
				}
				break;
			case FLOATS:
				enterOuterAlt(_localctx, 3);
				{
				setState(473);
				match(FLOATS);
				_localctx.t = environment.FLOAT
				}
				break;
			case BOOLS:
				enterOuterAlt(_localctx, 4);
				{
				setState(475);
				match(BOOLS);
				_localctx.t = environment.BOOLEAN 
				}
				break;
			case CHARACTERS:
				enterOuterAlt(_localctx, 5);
				{
				setState(477);
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
		case 22:
			return listParams_sempred((ListParamsContext)_localctx, predIndex);
		case 23:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean listParams_sempred(ListParamsContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return precpred(_ctx, 16);
		case 2:
			return precpred(_ctx, 15);
		case 3:
			return precpred(_ctx, 14);
		case 4:
			return precpred(_ctx, 13);
		case 5:
			return precpred(_ctx, 12);
		case 6:
			return precpred(_ctx, 11);
		case 7:
			return precpred(_ctx, 10);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3A\u01e4\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\3\2\3\2\3\2\3\2\3\3\6\3>\n\3\r\3\16\3?"+
		"\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4q\n\4\3\5\3\5\3\5"+
		"\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6\u008e\n\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3"+
		"\7\3\7\3\7\3\7\3\7\3\7\3\7\5\7\u009e\n\7\3\b\6\b\u00a1\n\b\r\b\16\b\u00a2"+
		"\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\5\t\u00c2\n\t\3\n\3\n\3"+
		"\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\5\n\u00d4\n\n\3"+
		"\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3"+
		"\13\3\13\5\13\u00e6\n\13\3\f\6\f\u00e9\n\f\r\f\16\f\u00ea\3\f\3\f\3\r"+
		"\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3\17"+
		"\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17"+
		"\3\17\3\17\3\17\5\17\u0110\n\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20"+
		"\3\20\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21"+
		"\3\21\5\21\u0129\n\21\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22"+
		"\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22"+
		"\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22"+
		"\3\22\3\22\3\22\5\22\u0154\n\22\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23"+
		"\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\25\3\25\3\25"+
		"\3\25\3\25\3\26\3\26\3\26\3\26\3\27\3\27\3\27\3\27\3\30\3\30\3\30\3\30"+
		"\3\30\3\30\3\30\3\30\3\30\7\30\u017f\n\30\f\30\16\30\u0182\13\30\3\31"+
		"\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31"+
		"\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\5\31\u019f"+
		"\n\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31"+
		"\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31"+
		"\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\31\7\31\u01c4\n\31\f\31\16\31\u01c7"+
		"\13\31\3\32\3\32\3\32\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33"+
		"\5\33\u01d6\n\33\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\5\34"+
		"\u01e2\n\34\3\34\2\4.\60\35\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \""+
		"$&(*,.\60\62\64\66\2\b\3\2\27\31\3\2/\60\3\2\61\63\4\2++--\4\2,,..\3\2"+
		"#$\2\u01fc\28\3\2\2\2\4=\3\2\2\2\6p\3\2\2\2\br\3\2\2\2\n\u008d\3\2\2\2"+
		"\f\u009d\3\2\2\2\16\u00a0\3\2\2\2\20\u00c1\3\2\2\2\22\u00d3\3\2\2\2\24"+
		"\u00e5\3\2\2\2\26\u00e8\3\2\2\2\30\u00ee\3\2\2\2\32\u00f3\3\2\2\2\34\u010f"+
		"\3\2\2\2\36\u0111\3\2\2\2 \u0128\3\2\2\2\"\u0153\3\2\2\2$\u0155\3\2\2"+
		"\2&\u015d\3\2\2\2(\u0164\3\2\2\2*\u016d\3\2\2\2,\u0171\3\2\2\2.\u0175"+
		"\3\2\2\2\60\u019e\3\2\2\2\62\u01c8\3\2\2\2\64\u01d5\3\2\2\2\66\u01e1\3"+
		"\2\2\289\5\4\3\29:\7\2\2\3:;\b\2\1\2;\3\3\2\2\2<>\5\6\4\2=<\3\2\2\2>?"+
		"\3\2\2\2?=\3\2\2\2?@\3\2\2\2@A\3\2\2\2AB\b\3\1\2B\5\3\2\2\2CD\5\b\5\2"+
		"DE\b\4\1\2Eq\3\2\2\2FG\5\n\6\2GH\b\4\1\2Hq\3\2\2\2IJ\5\f\7\2JK\b\4\1\2"+
		"Kq\3\2\2\2LM\5\30\r\2MN\b\4\1\2Nq\3\2\2\2OP\5\64\33\2PQ\b\4\1\2Qq\3\2"+
		"\2\2RS\5\20\t\2ST\b\4\1\2Tq\3\2\2\2UV\5\32\16\2VW\b\4\1\2Wq\3\2\2\2XY"+
		"\5\34\17\2YZ\b\4\1\2Zq\3\2\2\2[\\\5\22\n\2\\]\b\4\1\2]q\3\2\2\2^_\5\36"+
		"\20\2_`\b\4\1\2`q\3\2\2\2ab\5 \21\2bc\b\4\1\2cq\3\2\2\2de\5\"\22\2ef\b"+
		"\4\1\2fq\3\2\2\2gh\5$\23\2hi\b\4\1\2iq\3\2\2\2jk\5&\24\2kl\b\4\1\2lq\3"+
		"\2\2\2mn\5(\25\2no\b\4\1\2oq\3\2\2\2pC\3\2\2\2pF\3\2\2\2pI\3\2\2\2pL\3"+
		"\2\2\2pO\3\2\2\2pR\3\2\2\2pU\3\2\2\2pX\3\2\2\2p[\3\2\2\2p^\3\2\2\2pa\3"+
		"\2\2\2pd\3\2\2\2pg\3\2\2\2pj\3\2\2\2pm\3\2\2\2q\7\3\2\2\2rs\7\f\2\2st"+
		"\7\64\2\2tu\5\60\31\2uv\7\65\2\2vw\b\5\1\2w\t\3\2\2\2xy\7\b\2\2yz\7\""+
		"\2\2z{\78\2\2{|\5\66\34\2|}\7(\2\2}~\5\60\31\2~\177\b\6\1\2\177\u008e"+
		"\3\2\2\2\u0080\u0081\7\b\2\2\u0081\u0082\7\"\2\2\u0082\u0083\7(\2\2\u0083"+
		"\u0084\5\60\31\2\u0084\u0085\b\6\1\2\u0085\u008e\3\2\2\2\u0086\u0087\7"+
		"\b\2\2\u0087\u0088\7\"\2\2\u0088\u0089\78\2\2\u0089\u008a\5\66\34\2\u008a"+
		"\u008b\79\2\2\u008b\u008c\b\6\1\2\u008c\u008e\3\2\2\2\u008dx\3\2\2\2\u008d"+
		"\u0080\3\2\2\2\u008d\u0086\3\2\2\2\u008e\13\3\2\2\2\u008f\u0090\7\20\2"+
		"\2\u0090\u0091\7\"\2\2\u0091\u0092\78\2\2\u0092\u0093\5\66\34\2\u0093"+
		"\u0094\7(\2\2\u0094\u0095\5\60\31\2\u0095\u0096\b\7\1\2\u0096\u009e\3"+
		"\2\2\2\u0097\u0098\7\20\2\2\u0098\u0099\7\"\2\2\u0099\u009a\7(\2\2\u009a"+
		"\u009b\5\60\31\2\u009b\u009c\b\7\1\2\u009c\u009e\3\2\2\2\u009d\u008f\3"+
		"\2\2\2\u009d\u0097\3\2\2\2\u009e\r\3\2\2\2\u009f\u00a1\5\20\t\2\u00a0"+
		"\u009f\3\2\2\2\u00a1\u00a2\3\2\2\2\u00a2\u00a0\3\2\2\2\u00a2\u00a3\3\2"+
		"\2\2\u00a3\u00a4\3\2\2\2\u00a4\u00a5\b\b\1\2\u00a5\17\3\2\2\2\u00a6\u00a7"+
		"\7\r\2\2\u00a7\u00a8\5\60\31\2\u00a8\u00a9\7\66\2\2\u00a9\u00aa\5\4\3"+
		"\2\u00aa\u00ab\7\67\2\2\u00ab\u00ac\b\t\1\2\u00ac\u00c2\3\2\2\2\u00ad"+
		"\u00ae\7\r\2\2\u00ae\u00af\5\60\31\2\u00af\u00b0\7\66\2\2\u00b0\u00b1"+
		"\5\4\3\2\u00b1\u00b2\7\67\2\2\u00b2\u00b3\7\16\2\2\u00b3\u00b4\7\66\2"+
		"\2\u00b4\u00b5\5\4\3\2\u00b5\u00b6\7\67\2\2\u00b6\u00b7\b\t\1\2\u00b7"+
		"\u00c2\3\2\2\2\u00b8\u00b9\7\r\2\2\u00b9\u00ba\5\60\31\2\u00ba\u00bb\7"+
		"\66\2\2\u00bb\u00bc\5\4\3\2\u00bc\u00bd\7\67\2\2\u00bd\u00be\7\16\2\2"+
		"\u00be\u00bf\5\16\b\2\u00bf\u00c0\b\t\1\2\u00c0\u00c2\3\2\2\2\u00c1\u00a6"+
		"\3\2\2\2\u00c1\u00ad\3\2\2\2\u00c1\u00b8\3\2\2\2\u00c2\21\3\2\2\2\u00c3"+
		"\u00c4\7\23\2\2\u00c4\u00c5\5\60\31\2\u00c5\u00c6\7\66\2\2\u00c6\u00c7"+
		"\7\25\2\2\u00c7\u00c8\78\2\2\u00c8\u00c9\5\4\3\2\u00c9\u00ca\7\67\2\2"+
		"\u00ca\u00cb\b\n\1\2\u00cb\u00d4\3\2\2\2\u00cc\u00cd\7\23\2\2\u00cd\u00ce"+
		"\5\60\31\2\u00ce\u00cf\7\66\2\2\u00cf\u00d0\5\24\13\2\u00d0\u00d1\7\67"+
		"\2\2\u00d1\u00d2\b\n\1\2\u00d2\u00d4\3\2\2\2\u00d3\u00c3\3\2\2\2\u00d3"+
		"\u00cc\3\2\2\2\u00d4\23\3\2\2\2\u00d5\u00d6\7\24\2\2\u00d6\u00d7\5\60"+
		"\31\2\u00d7\u00d8\78\2\2\u00d8\u00d9\5\4\3\2\u00d9\u00da\5\26\f\2\u00da"+
		"\u00db\b\13\1\2\u00db\u00e6\3\2\2\2\u00dc\u00dd\7\24\2\2\u00dd\u00de\5"+
		"\60\31\2\u00de\u00df\78\2\2\u00df\u00e0\5\4\3\2\u00e0\u00e1\7\25\2\2\u00e1"+
		"\u00e2\78\2\2\u00e2\u00e3\5\4\3\2\u00e3\u00e4\b\13\1\2\u00e4\u00e6\3\2"+
		"\2\2\u00e5\u00d5\3\2\2\2\u00e5\u00dc\3\2\2\2\u00e6\25\3\2\2\2\u00e7\u00e9"+
		"\5\24\13\2\u00e8\u00e7\3\2\2\2\u00e9\u00ea\3\2\2\2\u00ea\u00e8\3\2\2\2"+
		"\u00ea\u00eb\3\2\2\2\u00eb\u00ec\3\2\2\2\u00ec\u00ed\b\f\1\2\u00ed\27"+
		"\3\2\2\2\u00ee\u00ef\7\"\2\2\u00ef\u00f0\7(\2\2\u00f0\u00f1\5\60\31\2"+
		"\u00f1\u00f2\b\r\1\2\u00f2\31\3\2\2\2\u00f3\u00f4\7\17\2\2\u00f4\u00f5"+
		"\5\60\31\2\u00f5\u00f6\7\66\2\2\u00f6\u00f7\5\4\3\2\u00f7\u00f8\7\67\2"+
		"\2\u00f8\u00f9\b\16\1\2\u00f9\33\3\2\2\2\u00fa\u00fb\7\21\2\2\u00fb\u00fc"+
		"\7\"\2\2\u00fc\u00fd\7\22\2\2\u00fd\u00fe\5\60\31\2\u00fe\u00ff\7\66\2"+
		"\2\u00ff\u0100\5\4\3\2\u0100\u0101\7\67\2\2\u0101\u0102\b\17\1\2\u0102"+
		"\u0110\3\2\2\2\u0103\u0104\7\21\2\2\u0104\u0105\7\"\2\2\u0105\u0106\7"+
		"\22\2\2\u0106\u0107\5\60\31\2\u0107\u0108\7:\2\2\u0108\u0109\7:\2\2\u0109"+
		"\u010a\5\60\31\2\u010a\u010b\7\66\2\2\u010b\u010c\5\4\3\2\u010c\u010d"+
		"\7\67\2\2\u010d\u010e\b\17\1\2\u010e\u0110\3\2\2\2\u010f\u00fa\3\2\2\2"+
		"\u010f\u0103\3\2\2\2\u0110\35\3\2\2\2\u0111\u0112\7\26\2\2\u0112\u0113"+
		"\5\60\31\2\u0113\u0114\7\16\2\2\u0114\u0115\7\66\2\2\u0115\u0116\5\4\3"+
		"\2\u0116\u0117\t\2\2\2\u0117\u0118\7\67\2\2\u0118\u0119\b\20\1\2\u0119"+
		"\37\3\2\2\2\u011a\u011b\7\30\2\2\u011b\u011c\7<\2\2\u011c\u0129\b\21\1"+
		"\2\u011d\u011e\7\30\2\2\u011e\u011f\5\60\31\2\u011f\u0120\7<\2\2\u0120"+
		"\u0121\b\21\1\2\u0121\u0129\3\2\2\2\u0122\u0123\7\27\2\2\u0123\u0124\7"+
		"<\2\2\u0124\u0129\b\21\1\2\u0125\u0126\7\31\2\2\u0126\u0127\7<\2\2\u0127"+
		"\u0129\b\21\1\2\u0128\u011a\3\2\2\2\u0128\u011d\3\2\2\2\u0128\u0122\3"+
		"\2\2\2\u0128\u0125\3\2\2\2\u0129!\3\2\2\2\u012a\u012b\7\b\2\2\u012b\u012c"+
		"\7\"\2\2\u012c\u012d\78\2\2\u012d\u012e\7=\2\2\u012e\u012f\5\66\34\2\u012f"+
		"\u0130\7>\2\2\u0130\u0131\7(\2\2\u0131\u0132\7=\2\2\u0132\u0133\5.\30"+
		"\2\u0133\u0134\7>\2\2\u0134\u0135\b\22\1\2\u0135\u0154\3\2\2\2\u0136\u0137"+
		"\7\b\2\2\u0137\u0138\7\"\2\2\u0138\u0139\78\2\2\u0139\u013a\7=\2\2\u013a"+
		"\u013b\5\66\34\2\u013b\u013c\7>\2\2\u013c\u013d\7(\2\2\u013d\u013e\7="+
		"\2\2\u013e\u013f\7>\2\2\u013f\u0140\b\22\1\2\u0140\u0154\3\2\2\2\u0141"+
		"\u0142\7\b\2\2\u0142\u0143\7\"\2\2\u0143\u0144\7(\2\2\u0144\u0145\7=\2"+
		"\2\u0145\u0146\5\66\34\2\u0146\u0147\7>\2\2\u0147\u0148\7\64\2\2\u0148"+
		"\u0149\7\65\2\2\u0149\u0154\3\2\2\2\u014a\u014b\7\b\2\2\u014b\u014c\7"+
		"\"\2\2\u014c\u014d\78\2\2\u014d\u014e\7=\2\2\u014e\u014f\5\66\34\2\u014f"+
		"\u0150\7>\2\2\u0150\u0151\7(\2\2\u0151\u0152\7\"\2\2\u0152\u0154\3\2\2"+
		"\2\u0153\u012a\3\2\2\2\u0153\u0136\3\2\2\2\u0153\u0141\3\2\2\2\u0153\u014a"+
		"\3\2\2\2\u0154#\3\2\2\2\u0155\u0156\7\"\2\2\u0156\u0157\7:\2\2\u0157\u0158"+
		"\7\32\2\2\u0158\u0159\7\64\2\2\u0159\u015a\5\60\31\2\u015a\u015b\7\65"+
		"\2\2\u015b\u015c\b\23\1\2\u015c%\3\2\2\2\u015d\u015e\7\"\2\2\u015e\u015f"+
		"\7:\2\2\u015f\u0160\7\33\2\2\u0160\u0161\7\64\2\2\u0161\u0162\7\65\2\2"+
		"\u0162\u0163\b\24\1\2\u0163\'\3\2\2\2\u0164\u0165\7\"\2\2\u0165\u0166"+
		"\7:\2\2\u0166\u0167\7\34\2\2\u0167\u0168\7\64\2\2\u0168\u0169\7\35\2\2"+
		"\u0169\u016a\78\2\2\u016a\u016b\5\60\31\2\u016b\u016c\7\65\2\2\u016c)"+
		"\3\2\2\2\u016d\u016e\7\"\2\2\u016e\u016f\7:\2\2\u016f\u0170\7\36\2\2\u0170"+
		"+\3\2\2\2\u0171\u0172\7\"\2\2\u0172\u0173\7:\2\2\u0173\u0174\7\37\2\2"+
		"\u0174-\3\2\2\2\u0175\u0176\b\30\1\2\u0176\u0177\5\60\31\2\u0177\u0178"+
		"\b\30\1\2\u0178\u0180\3\2\2\2\u0179\u017a\f\4\2\2\u017a\u017b\7;\2\2\u017b"+
		"\u017c\5\60\31\2\u017c\u017d\b\30\1\2\u017d\u017f\3\2\2\2\u017e\u0179"+
		"\3\2\2\2\u017f\u0182\3\2\2\2\u0180\u017e\3\2\2\2\u0180\u0181\3\2\2\2\u0181"+
		"/\3\2\2\2\u0182\u0180\3\2\2\2\u0183\u0184\b\31\1\2\u0184\u0185\7%\2\2"+
		"\u0185\u0186\5\60\31\13\u0186\u0187\b\31\1\2\u0187\u019f\3\2\2\2\u0188"+
		"\u0189\7\62\2\2\u0189\u018a\5\60\31\n\u018a\u018b\b\31\1\2\u018b\u019f"+
		"\3\2\2\2\u018c\u018d\7\64\2\2\u018d\u018e\5\60\31\2\u018e\u018f\7\65\2"+
		"\2\u018f\u0190\b\31\1\2\u0190\u019f\3\2\2\2\u0191\u0192\7 \2\2\u0192\u019f"+
		"\b\31\1\2\u0193\u0194\7!\2\2\u0194\u019f\b\31\1\2\u0195\u0196\7\n\2\2"+
		"\u0196\u019f\b\31\1\2\u0197\u0198\7\13\2\2\u0198\u019f\b\31\1\2\u0199"+
		"\u019a\7\t\2\2\u019a\u019f\b\31\1\2\u019b\u019c\5\62\32\2\u019c\u019d"+
		"\b\31\1\2\u019d\u019f\3\2\2\2\u019e\u0183\3\2\2\2\u019e\u0188\3\2\2\2"+
		"\u019e\u018c\3\2\2\2\u019e\u0191\3\2\2\2\u019e\u0193\3\2\2\2\u019e\u0195"+
		"\3\2\2\2\u019e\u0197\3\2\2\2\u019e\u0199\3\2\2\2\u019e\u019b\3\2\2\2\u019f"+
		"\u01c5\3\2\2\2\u01a0\u01a1\f\22\2\2\u01a1\u01a2\t\3\2\2\u01a2\u01a3\5"+
		"\60\31\23\u01a3\u01a4\b\31\1\2\u01a4\u01c4\3\2\2\2\u01a5\u01a6\f\21\2"+
		"\2\u01a6\u01a7\t\4\2\2\u01a7\u01a8\5\60\31\22\u01a8\u01a9\b\31\1\2\u01a9"+
		"\u01c4\3\2\2\2\u01aa\u01ab\f\20\2\2\u01ab\u01ac\t\5\2\2\u01ac\u01ad\5"+
		"\60\31\21\u01ad\u01ae\b\31\1\2\u01ae\u01c4\3\2\2\2\u01af\u01b0\f\17\2"+
		"\2\u01b0\u01b1\t\6\2\2\u01b1\u01b2\5\60\31\20\u01b2\u01b3\b\31\1\2\u01b3"+
		"\u01c4\3\2\2\2\u01b4\u01b5\f\16\2\2\u01b5\u01b6\t\7\2\2\u01b6\u01b7\5"+
		"\60\31\17\u01b7\u01b8\b\31\1\2\u01b8\u01c4\3\2\2\2\u01b9\u01ba\f\r\2\2"+
		"\u01ba\u01bb\7\'\2\2\u01bb\u01bc\5\60\31\16\u01bc\u01bd\b\31\1\2\u01bd"+
		"\u01c4\3\2\2\2\u01be\u01bf\f\f\2\2\u01bf\u01c0\7&\2\2\u01c0\u01c1\5\60"+
		"\31\r\u01c1\u01c2\b\31\1\2\u01c2\u01c4\3\2\2\2\u01c3\u01a0\3\2\2\2\u01c3"+
		"\u01a5\3\2\2\2\u01c3\u01aa\3\2\2\2\u01c3\u01af\3\2\2\2\u01c3\u01b4\3\2"+
		"\2\2\u01c3\u01b9\3\2\2\2\u01c3\u01be\3\2\2\2\u01c4\u01c7\3\2\2\2\u01c5"+
		"\u01c3\3\2\2\2\u01c5\u01c6\3\2\2\2\u01c6\61\3\2\2\2\u01c7\u01c5\3\2\2"+
		"\2\u01c8\u01c9\7\"\2\2\u01c9\u01ca\b\32\1\2\u01ca\63\3\2\2\2\u01cb\u01cc"+
		"\7\"\2\2\u01cc\u01cd\7)\2\2\u01cd\u01ce\5\60\31\2\u01ce\u01cf\b\33\1\2"+
		"\u01cf\u01d6\3\2\2\2\u01d0\u01d1\7\"\2\2\u01d1\u01d2\7*\2\2\u01d2\u01d3"+
		"\5\60\31\2\u01d3\u01d4\b\33\1\2\u01d4\u01d6\3\2\2\2\u01d5\u01cb\3\2\2"+
		"\2\u01d5\u01d0\3\2\2\2\u01d6\65\3\2\2\2\u01d7\u01d8\7\3\2\2\u01d8\u01e2"+
		"\b\34\1\2\u01d9\u01da\7\4\2\2\u01da\u01e2\b\34\1\2\u01db\u01dc\7\5\2\2"+
		"\u01dc\u01e2\b\34\1\2\u01dd\u01de\7\6\2\2\u01de\u01e2\b\34\1\2\u01df\u01e0"+
		"\7\7\2\2\u01e0\u01e2\b\34\1\2\u01e1\u01d7\3\2\2\2\u01e1\u01d9\3\2\2\2"+
		"\u01e1\u01db\3\2\2\2\u01e1\u01dd\3\2\2\2\u01e1\u01df\3\2\2\2\u01e2\67"+
		"\3\2\2\2\24?p\u008d\u009d\u00a2\u00c1\u00d3\u00e5\u00ea\u010f\u0128\u0153"+
		"\u0180\u019e\u01c3\u01c5\u01d5\u01e1";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
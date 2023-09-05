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
		REMOVELAST=25, REMOVE=26, AT=27, ISEMPTY=28, COUNT=29, FUNC=30, INOUT=31, 
		NUMBER=32, STRING=33, ID=34, DIF=35, IG_IG=36, NOT=37, OR=38, AND=39, 
		IG=40, IG_ADD=41, IG_SUB=42, MAY_IG=43, MEN_IG=44, MAYOR=45, MENOR=46, 
		MUL=47, DIV=48, ADD=49, SUB=50, PERCENT=51, PARIZQ=52, PARDER=53, LLAVEIZQ=54, 
		LLAVEDER=55, DOUBLEPTS=56, QUESTION=57, POINT=58, COMA=59, PTCOMA=60, 
		CORCHETEIZQ=61, CORCHETEDER=62, GUION_BAJO=63, AND_SIMPLE=64, WHITESPACE=65, 
		COMMENT=66, LINE_COMMENT=67;
	public static final int
		RULE_s = 0, RULE_block = 1, RULE_instruction = 2, RULE_printstmt = 3, 
		RULE_declarestmt = 4, RULE_constantstmt = 5, RULE_blockelifs = 6, RULE_ifstmt = 7, 
		RULE_switchstmt = 8, RULE_casestmt = 9, RULE_blockcases = 10, RULE_assignationstmt = 11, 
		RULE_whilestmt = 12, RULE_forstmt = 13, RULE_guardstmt = 14, RULE_transferstmt = 15, 
		RULE_declarevectorstmt = 16, RULE_accessfuncinstruction = 17, RULE_appendstmt = 18, 
		RULE_removelaststmt = 19, RULE_removestmt = 20, RULE_emptvecstmt = 21, 
		RULE_countvecstmt = 22, RULE_accessvecstmt = 23, RULE_assignationvecstmt = 24, 
		RULE_declarematrixstmt = 25, RULE_declarefuncstmt = 26, RULE_listParamsFunc = 27, 
		RULE_parameterfuncstmt = 28, RULE_expr = 29, RULE_intfunctionstmt = 30, 
		RULE_floatfunctionstmt = 31, RULE_stringfunctionstmt = 32, RULE_accessfuncstmt = 33, 
		RULE_accessstmt = 34, RULE_increaseanddecreasestmt = 35, RULE_type = 36, 
		RULE_listParams = 37;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "block", "instruction", "printstmt", "declarestmt", "constantstmt", 
			"blockelifs", "ifstmt", "switchstmt", "casestmt", "blockcases", "assignationstmt", 
			"whilestmt", "forstmt", "guardstmt", "transferstmt", "declarevectorstmt", 
			"accessfuncinstruction", "appendstmt", "removelaststmt", "removestmt", 
			"emptvecstmt", "countvecstmt", "accessvecstmt", "assignationvecstmt", 
			"declarematrixstmt", "declarefuncstmt", "listParamsFunc", "parameterfuncstmt", 
			"expr", "intfunctionstmt", "floatfunctionstmt", "stringfunctionstmt", 
			"accessfuncstmt", "accessstmt", "increaseanddecreasestmt", "type", "listParams"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'String'", "'Int'", "'Float'", "'Bool'", "'Character'", "'var'", 
			"'nil'", "'true'", "'false'", "'print'", "'if'", "'else'", "'while'", 
			"'let'", "'for'", "'in'", "'switch'", "'case'", "'default'", "'guard'", 
			"'continue'", "'return'", "'break'", "'append'", "'removeLast'", "'remove'", 
			"'at'", "'IsEmpty'", "'count'", "'func'", "'inout'", null, null, null, 
			"'!='", "'=='", "'!'", "'||'", "'&&'", "'='", "'+='", "'-='", "'>='", 
			"'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'-'", "'%'", "'('", "')'", 
			"'{'", "'}'", "':'", "'?'", "'.'", "','", "';'", "'['", "']'", "'_'", 
			"'&'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "STRINGS", "INTS", "FLOATS", "BOOLS", "CHARACTERS", "VAR", "NIL", 
			"TRUE", "FALSE", "PRINT", "IF", "ELSE", "WHILE", "LET", "FOR", "IN", 
			"SWITCH", "CASE", "DEFAULT", "GUARD", "CONTINUE", "RETURN", "BREAK", 
			"APPEND", "REMOVELAST", "REMOVE", "AT", "ISEMPTY", "COUNT", "FUNC", "INOUT", 
			"NUMBER", "STRING", "ID", "DIF", "IG_IG", "NOT", "OR", "AND", "IG", "IG_ADD", 
			"IG_SUB", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", "DIV", "ADD", 
			"SUB", "PERCENT", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "DOUBLEPTS", 
			"QUESTION", "POINT", "COMA", "PTCOMA", "CORCHETEIZQ", "CORCHETEDER", 
			"GUION_BAJO", "AND_SIMPLE", "WHITESPACE", "COMMENT", "LINE_COMMENT"
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
			setState(76);
			((SContext)_localctx).block = block();
			setState(77);
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
			setState(81); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(80);
					((BlockContext)_localctx).instruction = instruction();
					((BlockContext)_localctx).ins.add(((BlockContext)_localctx).instruction);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(83); 
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
		public AssignationvecstmtContext assignationvecstmt;
		public DeclarefuncstmtContext declarefuncstmt;
		public AccessfuncinstructionContext accessfuncinstruction;
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
		public AssignationvecstmtContext assignationvecstmt() {
			return getRuleContext(AssignationvecstmtContext.class,0);
		}
		public DeclarefuncstmtContext declarefuncstmt() {
			return getRuleContext(DeclarefuncstmtContext.class,0);
		}
		public AccessfuncinstructionContext accessfuncinstruction() {
			return getRuleContext(AccessfuncinstructionContext.class,0);
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
			setState(141);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(87);
				((InstructionContext)_localctx).printstmt = printstmt();
				 _localctx.inst = ((InstructionContext)_localctx).printstmt.prnt
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(90);
				((InstructionContext)_localctx).declarestmt = declarestmt();
				_localctx.inst = ((InstructionContext)_localctx).declarestmt.dec
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(93);
				((InstructionContext)_localctx).constantstmt = constantstmt();
				_localctx.inst = ((InstructionContext)_localctx).constantstmt.const
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(96);
				((InstructionContext)_localctx).assignationstmt = assignationstmt();
				_localctx.inst = ((InstructionContext)_localctx).assignationstmt.assign
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(99);
				((InstructionContext)_localctx).increaseanddecreasestmt = increaseanddecreasestmt();
				_localctx.inst = ((InstructionContext)_localctx).increaseanddecreasestmt.increasedecrease
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(102);
				((InstructionContext)_localctx).ifstmt = ifstmt();
				 _localctx.inst = ((InstructionContext)_localctx).ifstmt.ift 
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(105);
				((InstructionContext)_localctx).whilestmt = whilestmt();
				_localctx.inst = ((InstructionContext)_localctx).whilestmt.while
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(108);
				((InstructionContext)_localctx).forstmt = forstmt();
				_localctx.inst = ((InstructionContext)_localctx).forstmt.for
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(111);
				((InstructionContext)_localctx).switchstmt = switchstmt();
				_localctx.inst = ((InstructionContext)_localctx).switchstmt.switch
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(114);
				((InstructionContext)_localctx).guardstmt = guardstmt();
				_localctx.inst = ((InstructionContext)_localctx).guardstmt.gua
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(117);
				((InstructionContext)_localctx).transferstmt = transferstmt();
				_localctx.inst = ((InstructionContext)_localctx).transferstmt.tran
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(120);
				((InstructionContext)_localctx).declarevectorstmt = declarevectorstmt();
				_localctx.inst = ((InstructionContext)_localctx).declarevectorstmt.decvec
				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(123);
				((InstructionContext)_localctx).appendstmt = appendstmt();
				_localctx.inst = ((InstructionContext)_localctx).appendstmt.app
				}
				break;
			case 14:
				enterOuterAlt(_localctx, 14);
				{
				setState(126);
				((InstructionContext)_localctx).removelaststmt = removelaststmt();
				_localctx.inst = ((InstructionContext)_localctx).removelaststmt.removl
				}
				break;
			case 15:
				enterOuterAlt(_localctx, 15);
				{
				setState(129);
				((InstructionContext)_localctx).removestmt = removestmt();
				_localctx.inst = ((InstructionContext)_localctx).removestmt.remov
				}
				break;
			case 16:
				enterOuterAlt(_localctx, 16);
				{
				setState(132);
				((InstructionContext)_localctx).assignationvecstmt = assignationvecstmt();
				_localctx.inst = ((InstructionContext)_localctx).assignationvecstmt.assignvec
				}
				break;
			case 17:
				enterOuterAlt(_localctx, 17);
				{
				setState(135);
				((InstructionContext)_localctx).declarefuncstmt = declarefuncstmt();
				_localctx.inst = ((InstructionContext)_localctx).declarefuncstmt.decfunc
				}
				break;
			case 18:
				enterOuterAlt(_localctx, 18);
				{
				setState(138);
				((InstructionContext)_localctx).accessfuncinstruction = accessfuncinstruction();
				_localctx.inst = ((InstructionContext)_localctx).accessfuncinstruction.accessfuncin
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
		public ListParamsContext listParams;
		public TerminalNode PRINT() { return getToken(GrammarParser.PRINT, 0); }
		public TerminalNode PARIZQ() { return getToken(GrammarParser.PARIZQ, 0); }
		public ListParamsContext listParams() {
			return getRuleContext(ListParamsContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(GrammarParser.PARDER, 0); }
		public TerminalNode PTCOMA() { return getToken(GrammarParser.PTCOMA, 0); }
		public PrintstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_printstmt; }
	}

	public final PrintstmtContext printstmt() throws RecognitionException {
		PrintstmtContext _localctx = new PrintstmtContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_printstmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(143);
			((PrintstmtContext)_localctx).PRINT = match(PRINT);
			setState(144);
			match(PARIZQ);
			setState(145);
			((PrintstmtContext)_localctx).listParams = listParams(0);
			setState(146);
			match(PARDER);
			setState(148);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PTCOMA) {
				{
				setState(147);
				match(PTCOMA);
				}
			}

			 _localctx.prnt = instructions.NewPrint((((PrintstmtContext)_localctx).PRINT!=null?((PrintstmtContext)_localctx).PRINT.getLine():0), (((PrintstmtContext)_localctx).PRINT!=null?((PrintstmtContext)_localctx).PRINT.getCharPositionInLine():0), ((PrintstmtContext)_localctx).listParams.l)
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
		public TerminalNode PTCOMA() { return getToken(GrammarParser.PTCOMA, 0); }
		public TerminalNode QUESTION() { return getToken(GrammarParser.QUESTION, 0); }
		public DeclarestmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declarestmt; }
	}

	public final DeclarestmtContext declarestmt() throws RecognitionException {
		DeclarestmtContext _localctx = new DeclarestmtContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_declarestmt);
		int _la;
		try {
			setState(182);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(152);
				((DeclarestmtContext)_localctx).VAR = match(VAR);
				setState(153);
				((DeclarestmtContext)_localctx).ID = match(ID);
				setState(154);
				match(DOUBLEPTS);
				setState(155);
				((DeclarestmtContext)_localctx).type = type();
				setState(156);
				match(IG);
				setState(157);
				((DeclarestmtContext)_localctx).expr = expr(0);
				setState(159);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(158);
					match(PTCOMA);
					}
				}

				_localctx.dec = instructions.NewTodeclare((((DeclarestmtContext)_localctx).VAR!=null?((DeclarestmtContext)_localctx).VAR.getLine():0), (((DeclarestmtContext)_localctx).VAR!=null?((DeclarestmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarestmtContext)_localctx).ID!=null?((DeclarestmtContext)_localctx).ID.getText():null), ((DeclarestmtContext)_localctx).type.t, ((DeclarestmtContext)_localctx).expr.e, false)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(163);
				((DeclarestmtContext)_localctx).VAR = match(VAR);
				setState(164);
				((DeclarestmtContext)_localctx).ID = match(ID);
				setState(165);
				match(IG);
				setState(166);
				((DeclarestmtContext)_localctx).expr = expr(0);
				setState(168);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(167);
					match(PTCOMA);
					}
				}

				_localctx.dec = instructions.NewTodeclare((((DeclarestmtContext)_localctx).VAR!=null?((DeclarestmtContext)_localctx).VAR.getLine():0), (((DeclarestmtContext)_localctx).VAR!=null?((DeclarestmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarestmtContext)_localctx).ID!=null?((DeclarestmtContext)_localctx).ID.getText():null), environment.NULL, ((DeclarestmtContext)_localctx).expr.e, false)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(172);
				((DeclarestmtContext)_localctx).VAR = match(VAR);
				setState(173);
				((DeclarestmtContext)_localctx).ID = match(ID);
				setState(174);
				match(DOUBLEPTS);
				setState(175);
				((DeclarestmtContext)_localctx).type = type();
				setState(176);
				match(QUESTION);
				setState(178);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(177);
					match(PTCOMA);
					}
				}

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
		public TerminalNode PTCOMA() { return getToken(GrammarParser.PTCOMA, 0); }
		public ConstantstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_constantstmt; }
	}

	public final ConstantstmtContext constantstmt() throws RecognitionException {
		ConstantstmtContext _localctx = new ConstantstmtContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_constantstmt);
		int _la;
		try {
			setState(204);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(184);
				((ConstantstmtContext)_localctx).LET = match(LET);
				setState(185);
				((ConstantstmtContext)_localctx).ID = match(ID);
				setState(186);
				match(DOUBLEPTS);
				setState(187);
				((ConstantstmtContext)_localctx).type = type();
				setState(188);
				match(IG);
				setState(189);
				((ConstantstmtContext)_localctx).expr = expr(0);
				setState(191);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(190);
					match(PTCOMA);
					}
				}

				_localctx.const = instructions.NewTodeclare((((ConstantstmtContext)_localctx).LET!=null?((ConstantstmtContext)_localctx).LET.getLine():0), (((ConstantstmtContext)_localctx).LET!=null?((ConstantstmtContext)_localctx).LET.getCharPositionInLine():0), (((ConstantstmtContext)_localctx).ID!=null?((ConstantstmtContext)_localctx).ID.getText():null), ((ConstantstmtContext)_localctx).type.t, ((ConstantstmtContext)_localctx).expr.e, true)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(195);
				((ConstantstmtContext)_localctx).LET = match(LET);
				setState(196);
				((ConstantstmtContext)_localctx).ID = match(ID);
				setState(197);
				match(IG);
				setState(198);
				((ConstantstmtContext)_localctx).expr = expr(0);
				setState(200);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(199);
					match(PTCOMA);
					}
				}

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
			setState(207); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(206);
					((BlockelifsContext)_localctx).ifstmt = ifstmt();
					((BlockelifsContext)_localctx).elif.add(((BlockelifsContext)_localctx).ifstmt);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(209); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
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
			setState(240);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(213);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(214);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(215);
				match(LLAVEIZQ);
				setState(216);
				((IfstmtContext)_localctx).ifb = block();
				setState(217);
				match(LLAVEDER);
				 ((IfstmtContext)_localctx).ift =  instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).ifb.blk, nil); 
				                                        _localctx.el = ((IfstmtContext)_localctx).ifb.blk
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(220);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(221);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(222);
				match(LLAVEIZQ);
				setState(223);
				((IfstmtContext)_localctx).ifelseblck = block();
				setState(224);
				match(LLAVEDER);
				setState(225);
				match(ELSE);
				setState(226);
				match(LLAVEIZQ);
				setState(227);
				((IfstmtContext)_localctx).elseifblck = block();
				setState(228);
				match(LLAVEDER);
				 ((IfstmtContext)_localctx).ift =  instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).ifelseblck.blk, ((IfstmtContext)_localctx).elseifblck.blk); 
				                                                                                        ((IfstmtContext)_localctx).el =  ((IfstmtContext)_localctx).ifelseblck.blk;
				                                                                                        
				                                                                                        
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(231);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(232);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(233);
				match(LLAVEIZQ);
				setState(234);
				((IfstmtContext)_localctx).elif = block();
				setState(235);
				match(LLAVEDER);
				setState(236);
				match(ELSE);
				setState(237);
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
			setState(258);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(242);
				((SwitchstmtContext)_localctx).SWITCH = match(SWITCH);
				setState(243);
				((SwitchstmtContext)_localctx).expr = expr(0);
				setState(244);
				match(LLAVEIZQ);
				setState(245);
				match(DEFAULT);
				setState(246);
				match(DOUBLEPTS);
				setState(247);
				((SwitchstmtContext)_localctx).block = block();
				setState(248);
				match(LLAVEDER);
				 ((SwitchstmtContext)_localctx).switch =  instructions.NewSwitch((((SwitchstmtContext)_localctx).SWITCH!=null?((SwitchstmtContext)_localctx).SWITCH.getLine():0), (((SwitchstmtContext)_localctx).SWITCH!=null?((SwitchstmtContext)_localctx).SWITCH.getCharPositionInLine():0), ((SwitchstmtContext)_localctx).expr.e, nil,  nil, ((SwitchstmtContext)_localctx).block.blk);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(251);
				((SwitchstmtContext)_localctx).SWITCH = match(SWITCH);
				setState(252);
				((SwitchstmtContext)_localctx).expr = expr(0);
				setState(253);
				match(LLAVEIZQ);
				setState(254);
				((SwitchstmtContext)_localctx).casestmt = casestmt();
				setState(255);
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
			setState(276);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(260);
				((CasestmtContext)_localctx).CASE = match(CASE);
				setState(261);
				((CasestmtContext)_localctx).expr = expr(0);
				setState(262);
				match(DOUBLEPTS);
				setState(263);
				((CasestmtContext)_localctx).block = block();
				setState(264);
				((CasestmtContext)_localctx).blockcases = blockcases();
				 _localctx.cas = instructions.NewCase((((CasestmtContext)_localctx).CASE!=null?((CasestmtContext)_localctx).CASE.getLine():0),(((CasestmtContext)_localctx).CASE!=null?((CasestmtContext)_localctx).CASE.getCharPositionInLine():0), ((CasestmtContext)_localctx).expr.e,((CasestmtContext)_localctx).block.blk, ((CasestmtContext)_localctx).blockcases.blkcase)
				                                fmt.Println("entro")
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(267);
				((CasestmtContext)_localctx).CASE = match(CASE);
				setState(268);
				((CasestmtContext)_localctx).expr = expr(0);
				setState(269);
				match(DOUBLEPTS);
				setState(270);
				((CasestmtContext)_localctx).sen = block();
				setState(271);
				match(DEFAULT);
				setState(272);
				match(DOUBLEPTS);
				setState(273);
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
			setState(279); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(278);
					((BlockcasesContext)_localctx).casestmt = casestmt();
					((BlockcasesContext)_localctx).casedef.add(((BlockcasesContext)_localctx).casestmt);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(281); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,14,_ctx);
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
		public TerminalNode PTCOMA() { return getToken(GrammarParser.PTCOMA, 0); }
		public AssignationstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignationstmt; }
	}

	public final AssignationstmtContext assignationstmt() throws RecognitionException {
		AssignationstmtContext _localctx = new AssignationstmtContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_assignationstmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(285);
			((AssignationstmtContext)_localctx).ID = match(ID);
			setState(286);
			((AssignationstmtContext)_localctx).IG = match(IG);
			setState(287);
			((AssignationstmtContext)_localctx).expr = expr(0);
			setState(289);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PTCOMA) {
				{
				setState(288);
				match(PTCOMA);
				}
			}

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
			setState(293);
			((WhilestmtContext)_localctx).WHILE = match(WHILE);
			setState(294);
			((WhilestmtContext)_localctx).expr = expr(0);
			setState(295);
			match(LLAVEIZQ);
			setState(296);
			((WhilestmtContext)_localctx).block = block();
			setState(297);
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
			setState(321);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(300);
				((ForstmtContext)_localctx).FOR = match(FOR);
				setState(301);
				((ForstmtContext)_localctx).ID = match(ID);
				setState(302);
				match(IN);
				setState(303);
				((ForstmtContext)_localctx).expr = expr(0);
				setState(304);
				match(LLAVEIZQ);
				setState(305);
				((ForstmtContext)_localctx).block = block();
				setState(306);
				match(LLAVEDER);
				_localctx.for = instructions.NewFor((((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getLine():0), (((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getCharPositionInLine():0), (((ForstmtContext)_localctx).ID!=null?((ForstmtContext)_localctx).ID.getText():null), ((ForstmtContext)_localctx).expr.e, nil, ((ForstmtContext)_localctx).block.blk)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(309);
				((ForstmtContext)_localctx).FOR = match(FOR);
				setState(310);
				((ForstmtContext)_localctx).ID = match(ID);
				setState(311);
				match(IN);
				setState(312);
				((ForstmtContext)_localctx).first = expr(0);
				setState(313);
				match(POINT);
				setState(314);
				match(POINT);
				setState(315);
				((ForstmtContext)_localctx).second = expr(0);
				setState(316);
				match(LLAVEIZQ);
				setState(317);
				((ForstmtContext)_localctx).block = block();
				setState(318);
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
			setState(323);
			((GuardstmtContext)_localctx).GUARD = match(GUARD);
			setState(324);
			((GuardstmtContext)_localctx).expr = expr(0);
			setState(325);
			match(ELSE);
			setState(326);
			match(LLAVEIZQ);
			setState(327);
			((GuardstmtContext)_localctx).block = block();
			setState(328);
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
			setState(329);
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
		int _la;
		try {
			setState(354);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,21,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(332);
				((TransferstmtContext)_localctx).RETURN = match(RETURN);
				setState(334);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(333);
					match(PTCOMA);
					}
				}

				_localctx.tran = instructions.NewReturnIn((((TransferstmtContext)_localctx).RETURN!=null?((TransferstmtContext)_localctx).RETURN.getLine():0), (((TransferstmtContext)_localctx).RETURN!=null?((TransferstmtContext)_localctx).RETURN.getCharPositionInLine():0), nil)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(337);
				((TransferstmtContext)_localctx).RETURN = match(RETURN);
				setState(338);
				((TransferstmtContext)_localctx).expr = expr(0);
				setState(340);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(339);
					match(PTCOMA);
					}
				}

				_localctx.tran = instructions.NewReturnIn((((TransferstmtContext)_localctx).RETURN!=null?((TransferstmtContext)_localctx).RETURN.getLine():0), (((TransferstmtContext)_localctx).RETURN!=null?((TransferstmtContext)_localctx).RETURN.getCharPositionInLine():0), ((TransferstmtContext)_localctx).expr.e)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(344);
				((TransferstmtContext)_localctx).CONTINUE = match(CONTINUE);
				setState(346);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(345);
					match(PTCOMA);
					}
				}

				_localctx.tran = instructions.NewContinue((((TransferstmtContext)_localctx).CONTINUE!=null?((TransferstmtContext)_localctx).CONTINUE.getLine():0), (((TransferstmtContext)_localctx).CONTINUE!=null?((TransferstmtContext)_localctx).CONTINUE.getCharPositionInLine():0))
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(349);
				((TransferstmtContext)_localctx).BREAK = match(BREAK);
				setState(351);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(350);
					match(PTCOMA);
					}
				}

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
		public Token F;
		public Token S;
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
		public TerminalNode PTCOMA() { return getToken(GrammarParser.PTCOMA, 0); }
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
		int _la;
		try {
			setState(409);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,26,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(356);
				((DeclarevectorstmtContext)_localctx).VAR = match(VAR);
				setState(357);
				((DeclarevectorstmtContext)_localctx).ID = match(ID);
				setState(358);
				match(DOUBLEPTS);
				setState(359);
				match(CORCHETEIZQ);
				setState(360);
				((DeclarevectorstmtContext)_localctx).type = type();
				setState(361);
				match(CORCHETEDER);
				setState(362);
				match(IG);
				setState(363);
				match(CORCHETEIZQ);
				setState(364);
				((DeclarevectorstmtContext)_localctx).listParams = listParams(0);
				setState(365);
				match(CORCHETEDER);
				setState(367);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(366);
					match(PTCOMA);
					}
				}

				_localctx.decvec = instructions.NewToDeclareVector((((DeclarevectorstmtContext)_localctx).VAR!=null?((DeclarevectorstmtContext)_localctx).VAR.getLine():0), (((DeclarevectorstmtContext)_localctx).VAR!=null?((DeclarevectorstmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarevectorstmtContext)_localctx).ID!=null?((DeclarevectorstmtContext)_localctx).ID.getText():null), ((DeclarevectorstmtContext)_localctx).type.t, ((DeclarevectorstmtContext)_localctx).listParams.l, "")
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(371);
				((DeclarevectorstmtContext)_localctx).VAR = match(VAR);
				setState(372);
				((DeclarevectorstmtContext)_localctx).ID = match(ID);
				setState(373);
				match(DOUBLEPTS);
				setState(374);
				match(CORCHETEIZQ);
				setState(375);
				((DeclarevectorstmtContext)_localctx).type = type();
				setState(376);
				match(CORCHETEDER);
				setState(377);
				match(IG);
				setState(378);
				match(CORCHETEIZQ);
				setState(379);
				match(CORCHETEDER);
				setState(381);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(380);
					match(PTCOMA);
					}
				}

				_localctx.decvec = instructions.NewToDeclareVector((((DeclarevectorstmtContext)_localctx).VAR!=null?((DeclarevectorstmtContext)_localctx).VAR.getLine():0), (((DeclarevectorstmtContext)_localctx).VAR!=null?((DeclarevectorstmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarevectorstmtContext)_localctx).ID!=null?((DeclarevectorstmtContext)_localctx).ID.getText():null), ((DeclarevectorstmtContext)_localctx).type.t, nil, "")
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(385);
				match(VAR);
				setState(386);
				match(ID);
				setState(387);
				match(IG);
				setState(388);
				match(CORCHETEIZQ);
				setState(389);
				type();
				setState(390);
				match(CORCHETEDER);
				setState(391);
				match(PARIZQ);
				setState(392);
				match(PARDER);
				setState(394);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(393);
					match(PTCOMA);
					}
				}

				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(396);
				((DeclarevectorstmtContext)_localctx).VAR = match(VAR);
				setState(397);
				((DeclarevectorstmtContext)_localctx).F = match(ID);
				setState(398);
				match(DOUBLEPTS);
				setState(399);
				match(CORCHETEIZQ);
				setState(400);
				((DeclarevectorstmtContext)_localctx).type = type();
				setState(401);
				match(CORCHETEDER);
				setState(402);
				match(IG);
				setState(403);
				((DeclarevectorstmtContext)_localctx).S = match(ID);
				setState(405);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(404);
					match(PTCOMA);
					}
				}

				_localctx.decvec = instructions.NewToDeclareVector((((DeclarevectorstmtContext)_localctx).VAR!=null?((DeclarevectorstmtContext)_localctx).VAR.getLine():0), (((DeclarevectorstmtContext)_localctx).VAR!=null?((DeclarevectorstmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarevectorstmtContext)_localctx).F!=null?((DeclarevectorstmtContext)_localctx).F.getText():null), ((DeclarevectorstmtContext)_localctx).type.t, nil, (((DeclarevectorstmtContext)_localctx).S!=null?((DeclarevectorstmtContext)_localctx).S.getText():null))
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

	public static class AccessfuncinstructionContext extends ParserRuleContext {
		public interfaces.Instruction accessfuncin;
		public Token ID;
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode PARIZQ() { return getToken(GrammarParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(GrammarParser.PARDER, 0); }
		public AccessfuncinstructionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_accessfuncinstruction; }
	}

	public final AccessfuncinstructionContext accessfuncinstruction() throws RecognitionException {
		AccessfuncinstructionContext _localctx = new AccessfuncinstructionContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_accessfuncinstruction);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(411);
			((AccessfuncinstructionContext)_localctx).ID = match(ID);
			setState(412);
			match(PARIZQ);
			setState(413);
			match(PARDER);
			_localctx.accessfuncin = instructions.NewCallFunction((((AccessfuncinstructionContext)_localctx).ID!=null?((AccessfuncinstructionContext)_localctx).ID.getLine():0), (((AccessfuncinstructionContext)_localctx).ID!=null?((AccessfuncinstructionContext)_localctx).ID.getCharPositionInLine():0), (((AccessfuncinstructionContext)_localctx).ID!=null?((AccessfuncinstructionContext)_localctx).ID.getText():null))
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
		public TerminalNode PTCOMA() { return getToken(GrammarParser.PTCOMA, 0); }
		public AppendstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_appendstmt; }
	}

	public final AppendstmtContext appendstmt() throws RecognitionException {
		AppendstmtContext _localctx = new AppendstmtContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_appendstmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(416);
			((AppendstmtContext)_localctx).ID = match(ID);
			setState(417);
			match(POINT);
			setState(418);
			match(APPEND);
			setState(419);
			match(PARIZQ);
			setState(420);
			((AppendstmtContext)_localctx).expr = expr(0);
			setState(421);
			match(PARDER);
			setState(423);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PTCOMA) {
				{
				setState(422);
				match(PTCOMA);
				}
			}

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
		public TerminalNode PTCOMA() { return getToken(GrammarParser.PTCOMA, 0); }
		public RemovelaststmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_removelaststmt; }
	}

	public final RemovelaststmtContext removelaststmt() throws RecognitionException {
		RemovelaststmtContext _localctx = new RemovelaststmtContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_removelaststmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(427);
			((RemovelaststmtContext)_localctx).ID = match(ID);
			setState(428);
			match(POINT);
			setState(429);
			match(REMOVELAST);
			setState(430);
			match(PARIZQ);
			setState(431);
			match(PARDER);
			setState(433);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PTCOMA) {
				{
				setState(432);
				match(PTCOMA);
				}
			}

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
		public Token ID;
		public ExprContext expr;
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
		public TerminalNode PTCOMA() { return getToken(GrammarParser.PTCOMA, 0); }
		public RemovestmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_removestmt; }
	}

	public final RemovestmtContext removestmt() throws RecognitionException {
		RemovestmtContext _localctx = new RemovestmtContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_removestmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(437);
			((RemovestmtContext)_localctx).ID = match(ID);
			setState(438);
			match(POINT);
			setState(439);
			match(REMOVE);
			setState(440);
			match(PARIZQ);
			setState(441);
			match(AT);
			setState(442);
			match(DOUBLEPTS);
			setState(443);
			((RemovestmtContext)_localctx).expr = expr(0);
			setState(444);
			match(PARDER);
			setState(446);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PTCOMA) {
				{
				setState(445);
				match(PTCOMA);
				}
			}

			_localctx.remov = instructions.NewRemove((((RemovestmtContext)_localctx).ID!=null?((RemovestmtContext)_localctx).ID.getLine():0), (((RemovestmtContext)_localctx).ID!=null?((RemovestmtContext)_localctx).ID.getCharPositionInLine():0), (((RemovestmtContext)_localctx).ID!=null?((RemovestmtContext)_localctx).ID.getText():null), ((RemovestmtContext)_localctx).expr.e)
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
		public Token ID;
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
		enterRule(_localctx, 42, RULE_emptvecstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(450);
			((EmptvecstmtContext)_localctx).ID = match(ID);
			setState(451);
			match(POINT);
			setState(452);
			match(ISEMPTY);
			_localctx.emptyvec = expressions.NewEmptyVector((((EmptvecstmtContext)_localctx).ID!=null?((EmptvecstmtContext)_localctx).ID.getLine():0), (((EmptvecstmtContext)_localctx).ID!=null?((EmptvecstmtContext)_localctx).ID.getCharPositionInLine():0), (((EmptvecstmtContext)_localctx).ID!=null?((EmptvecstmtContext)_localctx).ID.getText():null))
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
		public Token ID;
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
		enterRule(_localctx, 44, RULE_countvecstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(455);
			((CountvecstmtContext)_localctx).ID = match(ID);
			setState(456);
			match(POINT);
			setState(457);
			match(COUNT);
			_localctx.count = expressions.NewCount((((CountvecstmtContext)_localctx).ID!=null?((CountvecstmtContext)_localctx).ID.getLine():0), (((CountvecstmtContext)_localctx).ID!=null?((CountvecstmtContext)_localctx).ID.getCharPositionInLine():0), (((CountvecstmtContext)_localctx).ID!=null?((CountvecstmtContext)_localctx).ID.getText():null))
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

	public static class AccessvecstmtContext extends ParserRuleContext {
		public interfaces.Expression accessvec;
		public Token ID;
		public ExprContext expr;
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode CORCHETEIZQ() { return getToken(GrammarParser.CORCHETEIZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode CORCHETEDER() { return getToken(GrammarParser.CORCHETEDER, 0); }
		public AccessvecstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_accessvecstmt; }
	}

	public final AccessvecstmtContext accessvecstmt() throws RecognitionException {
		AccessvecstmtContext _localctx = new AccessvecstmtContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_accessvecstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(460);
			((AccessvecstmtContext)_localctx).ID = match(ID);
			setState(461);
			match(CORCHETEIZQ);
			setState(462);
			((AccessvecstmtContext)_localctx).expr = expr(0);
			setState(463);
			match(CORCHETEDER);
			_localctx.accessvec = expressions.NewAccessVector((((AccessvecstmtContext)_localctx).ID!=null?((AccessvecstmtContext)_localctx).ID.getLine():0), (((AccessvecstmtContext)_localctx).ID!=null?((AccessvecstmtContext)_localctx).ID.getCharPositionInLine():0), (((AccessvecstmtContext)_localctx).ID!=null?((AccessvecstmtContext)_localctx).ID.getText():null), ((AccessvecstmtContext)_localctx).expr.e)
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

	public static class AssignationvecstmtContext extends ParserRuleContext {
		public interfaces.Instruction assignvec;
		public Token ID;
		public ExprContext expprim;
		public ExprContext expsegundo;
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode CORCHETEIZQ() { return getToken(GrammarParser.CORCHETEIZQ, 0); }
		public TerminalNode CORCHETEDER() { return getToken(GrammarParser.CORCHETEDER, 0); }
		public TerminalNode IG() { return getToken(GrammarParser.IG, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode PTCOMA() { return getToken(GrammarParser.PTCOMA, 0); }
		public AssignationvecstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignationvecstmt; }
	}

	public final AssignationvecstmtContext assignationvecstmt() throws RecognitionException {
		AssignationvecstmtContext _localctx = new AssignationvecstmtContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_assignationvecstmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(466);
			((AssignationvecstmtContext)_localctx).ID = match(ID);
			setState(467);
			match(CORCHETEIZQ);
			setState(468);
			((AssignationvecstmtContext)_localctx).expprim = expr(0);
			setState(469);
			match(CORCHETEDER);
			setState(470);
			match(IG);
			setState(471);
			((AssignationvecstmtContext)_localctx).expsegundo = expr(0);
			setState(473);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PTCOMA) {
				{
				setState(472);
				match(PTCOMA);
				}
			}

			_localctx.assignvec = instructions.NewAssignationVector((((AssignationvecstmtContext)_localctx).ID!=null?((AssignationvecstmtContext)_localctx).ID.getLine():0), (((AssignationvecstmtContext)_localctx).ID!=null?((AssignationvecstmtContext)_localctx).ID.getCharPositionInLine():0), (((AssignationvecstmtContext)_localctx).ID!=null?((AssignationvecstmtContext)_localctx).ID.getText():null), ((AssignationvecstmtContext)_localctx).expprim.e, ((AssignationvecstmtContext)_localctx).expsegundo.e)
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

	public static class DeclarematrixstmtContext extends ParserRuleContext {
		public interfaces.Instruction decmatrix;
		public TerminalNode VAR() { return getToken(GrammarParser.VAR, 0); }
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode PARIZQ() { return getToken(GrammarParser.PARIZQ, 0); }
		public TerminalNode DOUBLEPTS() { return getToken(GrammarParser.DOUBLEPTS, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(GrammarParser.PARDER, 0); }
		public TerminalNode IG() { return getToken(GrammarParser.IG, 0); }
		public TerminalNode PTCOMA() { return getToken(GrammarParser.PTCOMA, 0); }
		public DeclarematrixstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declarematrixstmt; }
	}

	public final DeclarematrixstmtContext declarematrixstmt() throws RecognitionException {
		DeclarematrixstmtContext _localctx = new DeclarematrixstmtContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_declarematrixstmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(477);
			match(VAR);
			setState(478);
			match(ID);
			setState(479);
			match(PARIZQ);
			setState(480);
			match(DOUBLEPTS);
			setState(481);
			type();
			setState(482);
			match(PARDER);
			setState(483);
			match(IG);
			setState(485);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PTCOMA) {
				{
				setState(484);
				match(PTCOMA);
				}
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

	public static class DeclarefuncstmtContext extends ParserRuleContext {
		public interfaces.Instruction decfunc;
		public Token ID;
		public ListParamsContext listParams;
		public TypeContext type;
		public BlockContext block;
		public TerminalNode FUNC() { return getToken(GrammarParser.FUNC, 0); }
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode PARIZQ() { return getToken(GrammarParser.PARIZQ, 0); }
		public ListParamsContext listParams() {
			return getRuleContext(ListParamsContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(GrammarParser.PARDER, 0); }
		public TerminalNode SUB() { return getToken(GrammarParser.SUB, 0); }
		public TerminalNode MAYOR() { return getToken(GrammarParser.MAYOR, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(GrammarParser.LLAVEIZQ, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(GrammarParser.LLAVEDER, 0); }
		public DeclarefuncstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declarefuncstmt; }
	}

	public final DeclarefuncstmtContext declarefuncstmt() throws RecognitionException {
		DeclarefuncstmtContext _localctx = new DeclarefuncstmtContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_declarefuncstmt);
		try {
			setState(531);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,32,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(487);
				match(FUNC);
				setState(488);
				((DeclarefuncstmtContext)_localctx).ID = match(ID);
				setState(489);
				match(PARIZQ);
				setState(490);
				((DeclarefuncstmtContext)_localctx).listParams = listParams(0);
				setState(491);
				match(PARDER);
				setState(492);
				match(SUB);
				setState(493);
				match(MAYOR);
				setState(494);
				((DeclarefuncstmtContext)_localctx).type = type();
				setState(495);
				match(LLAVEIZQ);
				setState(496);
				((DeclarefuncstmtContext)_localctx).block = block();
				setState(497);
				match(LLAVEDER);
				_localctx.decfunc = instructions.NewToDeclareFunction((((DeclarefuncstmtContext)_localctx).ID!=null?((DeclarefuncstmtContext)_localctx).ID.getLine():0), (((DeclarefuncstmtContext)_localctx).ID!=null?((DeclarefuncstmtContext)_localctx).ID.getCharPositionInLine():0), (((DeclarefuncstmtContext)_localctx).ID!=null?((DeclarefuncstmtContext)_localctx).ID.getText():null), ((DeclarefuncstmtContext)_localctx).listParams.l, ((DeclarefuncstmtContext)_localctx).type.t, ((DeclarefuncstmtContext)_localctx).block.blk, 1)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(500);
				match(FUNC);
				setState(501);
				((DeclarefuncstmtContext)_localctx).ID = match(ID);
				setState(502);
				match(PARIZQ);
				setState(503);
				((DeclarefuncstmtContext)_localctx).listParams = listParams(0);
				setState(504);
				match(PARDER);
				setState(505);
				match(LLAVEIZQ);
				setState(506);
				((DeclarefuncstmtContext)_localctx).block = block();
				setState(507);
				match(LLAVEDER);
				_localctx.decfunc = instructions.NewToDeclareFunction((((DeclarefuncstmtContext)_localctx).ID!=null?((DeclarefuncstmtContext)_localctx).ID.getLine():0), (((DeclarefuncstmtContext)_localctx).ID!=null?((DeclarefuncstmtContext)_localctx).ID.getCharPositionInLine():0), (((DeclarefuncstmtContext)_localctx).ID!=null?((DeclarefuncstmtContext)_localctx).ID.getText():null), ((DeclarefuncstmtContext)_localctx).listParams.l, environment.NULL, ((DeclarefuncstmtContext)_localctx).block.blk, 2)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(510);
				match(FUNC);
				setState(511);
				((DeclarefuncstmtContext)_localctx).ID = match(ID);
				setState(512);
				match(PARIZQ);
				setState(513);
				match(PARDER);
				setState(514);
				match(SUB);
				setState(515);
				match(MAYOR);
				setState(516);
				((DeclarefuncstmtContext)_localctx).type = type();
				setState(517);
				match(LLAVEIZQ);
				setState(518);
				((DeclarefuncstmtContext)_localctx).block = block();
				setState(519);
				match(LLAVEDER);
				_localctx.decfunc = instructions.NewToDeclareFunction((((DeclarefuncstmtContext)_localctx).ID!=null?((DeclarefuncstmtContext)_localctx).ID.getLine():0), (((DeclarefuncstmtContext)_localctx).ID!=null?((DeclarefuncstmtContext)_localctx).ID.getCharPositionInLine():0), (((DeclarefuncstmtContext)_localctx).ID!=null?((DeclarefuncstmtContext)_localctx).ID.getText():null), nil, ((DeclarefuncstmtContext)_localctx).type.t, ((DeclarefuncstmtContext)_localctx).block.blk, 3)
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(522);
				match(FUNC);
				setState(523);
				((DeclarefuncstmtContext)_localctx).ID = match(ID);
				setState(524);
				match(PARIZQ);
				setState(525);
				match(PARDER);
				setState(526);
				match(LLAVEIZQ);
				setState(527);
				((DeclarefuncstmtContext)_localctx).block = block();
				setState(528);
				match(LLAVEDER);

				    if(((DeclarefuncstmtContext)_localctx).block.blk!=nil){ 
				         _localctx.decfunc = instructions.NewToDeclareFunction((((DeclarefuncstmtContext)_localctx).ID!=null?((DeclarefuncstmtContext)_localctx).ID.getLine():0), (((DeclarefuncstmtContext)_localctx).ID!=null?((DeclarefuncstmtContext)_localctx).ID.getCharPositionInLine():0), (((DeclarefuncstmtContext)_localctx).ID!=null?((DeclarefuncstmtContext)_localctx).ID.getText():null), nil, environment.NULL, ((DeclarefuncstmtContext)_localctx).block.blk, 4)
				    }

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

	public static class ListParamsFuncContext extends ParserRuleContext {
		public []interface{} lf;
		public ListParamsFuncContext listf;
		public ParameterfuncstmtContext parameterfuncstmt;
		public ParameterfuncstmtContext parameterfuncstmt() {
			return getRuleContext(ParameterfuncstmtContext.class,0);
		}
		public TerminalNode COMA() { return getToken(GrammarParser.COMA, 0); }
		public ListParamsFuncContext listParamsFunc() {
			return getRuleContext(ListParamsFuncContext.class,0);
		}
		public ListParamsFuncContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listParamsFunc; }
	}

	public final ListParamsFuncContext listParamsFunc() throws RecognitionException {
		return listParamsFunc(0);
	}

	private ListParamsFuncContext listParamsFunc(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ListParamsFuncContext _localctx = new ListParamsFuncContext(_ctx, _parentState);
		ListParamsFuncContext _prevctx = _localctx;
		int _startState = 54;
		enterRecursionRule(_localctx, 54, RULE_listParamsFunc, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(534);
			((ListParamsFuncContext)_localctx).parameterfuncstmt = parameterfuncstmt();

			            _localctx.lf = []interface{}{}
			            _localctx.lf = append(_localctx.lf, ((ListParamsFuncContext)_localctx).parameterfuncstmt.parameterfunc)
			        
			}
			_ctx.stop = _input.LT(-1);
			setState(542);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,33,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ListParamsFuncContext(_parentctx, _parentState);
					_localctx.listf = _prevctx;
					_localctx.listf = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_listParamsFunc);
					setState(537);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(538);
					match(COMA);

					                                          var arrf []interface{}
					                                          arrf = append(((ListParamsFuncContext)_localctx).listf.lf, ((ListParamsFuncContext)_localctx).parameterfuncstmt.parameterfunc)
					                                          _localctx.lf = arrf
					                                      
					}
					} 
				}
				setState(544);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,33,_ctx);
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

	public static class ParameterfuncstmtContext extends ParserRuleContext {
		public interfaces.Expression parameterfunc;
		public Token exte;
		public List<TerminalNode> ID() { return getTokens(GrammarParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(GrammarParser.ID, i);
		}
		public TerminalNode DOUBLEPTS() { return getToken(GrammarParser.DOUBLEPTS, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode INOUT() { return getToken(GrammarParser.INOUT, 0); }
		public TerminalNode GUION_BAJO() { return getToken(GrammarParser.GUION_BAJO, 0); }
		public ParameterfuncstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parameterfuncstmt; }
	}

	public final ParameterfuncstmtContext parameterfuncstmt() throws RecognitionException {
		ParameterfuncstmtContext _localctx = new ParameterfuncstmtContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_parameterfuncstmt);
		int _la;
		try {
			setState(558);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,36,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(545);
				match(ID);
				setState(546);
				match(DOUBLEPTS);
				setState(548);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==INOUT) {
					{
					setState(547);
					match(INOUT);
					}
				}

				setState(550);
				type();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(551);
				((ParameterfuncstmtContext)_localctx).exte = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==ID || _la==GUION_BAJO) ) {
					((ParameterfuncstmtContext)_localctx).exte = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(552);
				match(ID);
				setState(553);
				match(DOUBLEPTS);
				setState(555);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==INOUT) {
					{
					setState(554);
					match(INOUT);
					}
				}

				setState(557);
				type();
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
		public Token NIL;
		public AccessstmtContext accessstmt;
		public EmptvecstmtContext emptvecstmt;
		public CountvecstmtContext countvecstmt;
		public AccessvecstmtContext accessvecstmt;
		public IntfunctionstmtContext intfunctionstmt;
		public FloatfunctionstmtContext floatfunctionstmt;
		public StringfunctionstmtContext stringfunctionstmt;
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
		public EmptvecstmtContext emptvecstmt() {
			return getRuleContext(EmptvecstmtContext.class,0);
		}
		public CountvecstmtContext countvecstmt() {
			return getRuleContext(CountvecstmtContext.class,0);
		}
		public AccessvecstmtContext accessvecstmt() {
			return getRuleContext(AccessvecstmtContext.class,0);
		}
		public IntfunctionstmtContext intfunctionstmt() {
			return getRuleContext(IntfunctionstmtContext.class,0);
		}
		public FloatfunctionstmtContext floatfunctionstmt() {
			return getRuleContext(FloatfunctionstmtContext.class,0);
		}
		public StringfunctionstmtContext stringfunctionstmt() {
			return getRuleContext(StringfunctionstmtContext.class,0);
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
		int _startState = 58;
		enterRecursionRule(_localctx, 58, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(605);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,37,_ctx) ) {
			case 1:
				{
				setState(561);
				((ExprContext)_localctx).op = match(NOT);
				setState(562);
				((ExprContext)_localctx).left = ((ExprContext)_localctx).expr = expr(15);
				 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).left.e) 
				}
				break;
			case 2:
				{
				setState(565);
				((ExprContext)_localctx).op = match(SUB);
				setState(566);
				((ExprContext)_localctx).left = ((ExprContext)_localctx).expr = expr(14);
				 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e,  (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), nil)
				}
				break;
			case 3:
				{
				setState(569);
				match(PARIZQ);
				setState(570);
				((ExprContext)_localctx).expr = expr(0);
				setState(571);
				match(PARDER);
				 _localctx.e = ((ExprContext)_localctx).expr.e 
				}
				break;
			case 4:
				{
				setState(574);
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
			case 5:
				{
				setState(576);
				((ExprContext)_localctx).STRING = match(STRING);

				        str := (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getText():null)
				        _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getLine():0), (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getCharPositionInLine():0), str[1:len(str)-1],environment.STRING)
				    
				}
				break;
			case 6:
				{
				setState(578);
				((ExprContext)_localctx).TRUE = match(TRUE);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).TRUE!=null?((ExprContext)_localctx).TRUE.getLine():0), (((ExprContext)_localctx).TRUE!=null?((ExprContext)_localctx).TRUE.getCharPositionInLine():0), true, environment.BOOLEAN) 
				}
				break;
			case 7:
				{
				setState(580);
				((ExprContext)_localctx).FALSE = match(FALSE);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).FALSE!=null?((ExprContext)_localctx).FALSE.getLine():0), (((ExprContext)_localctx).FALSE!=null?((ExprContext)_localctx).FALSE.getCharPositionInLine():0), false, environment.BOOLEAN) 
				}
				break;
			case 8:
				{
				setState(582);
				((ExprContext)_localctx).NIL = match(NIL);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).NIL!=null?((ExprContext)_localctx).NIL.getLine():0), (((ExprContext)_localctx).NIL!=null?((ExprContext)_localctx).NIL.getCharPositionInLine():0), nil, environment.NULL) 
				}
				break;
			case 9:
				{
				setState(584);
				((ExprContext)_localctx).accessstmt = accessstmt();
				_localctx.e = ((ExprContext)_localctx).accessstmt.access
				}
				break;
			case 10:
				{
				setState(587);
				((ExprContext)_localctx).emptvecstmt = emptvecstmt();
				_localctx.e = ((ExprContext)_localctx).emptvecstmt.emptyvec
				}
				break;
			case 11:
				{
				setState(590);
				((ExprContext)_localctx).countvecstmt = countvecstmt();
				_localctx.e = ((ExprContext)_localctx).countvecstmt.count
				}
				break;
			case 12:
				{
				setState(593);
				((ExprContext)_localctx).accessvecstmt = accessvecstmt();
				_localctx.e = ((ExprContext)_localctx).accessvecstmt.accessvec
				}
				break;
			case 13:
				{
				setState(596);
				((ExprContext)_localctx).intfunctionstmt = intfunctionstmt();
				_localctx.e = ((ExprContext)_localctx).intfunctionstmt.intfunc
				}
				break;
			case 14:
				{
				setState(599);
				((ExprContext)_localctx).floatfunctionstmt = floatfunctionstmt();
				_localctx.e = ((ExprContext)_localctx).floatfunctionstmt.floatfunc
				}
				break;
			case 15:
				{
				setState(602);
				((ExprContext)_localctx).stringfunctionstmt = stringfunctionstmt();
				_localctx.e = ((ExprContext)_localctx).stringfunctionstmt.stringfunc
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(644);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,39,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(642);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,38,_ctx) ) {
					case 1:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(607);
						if (!(precpred(_ctx, 22))) throw new FailedPredicateException(this, "precpred(_ctx, 22)");
						setState(608);
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
						setState(609);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(23);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 2:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(612);
						if (!(precpred(_ctx, 21))) throw new FailedPredicateException(this, "precpred(_ctx, 21)");
						setState(613);
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
						setState(614);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(22);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 3:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(617);
						if (!(precpred(_ctx, 20))) throw new FailedPredicateException(this, "precpred(_ctx, 20)");
						setState(618);
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
						setState(619);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(21);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 4:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(622);
						if (!(precpred(_ctx, 19))) throw new FailedPredicateException(this, "precpred(_ctx, 19)");
						setState(623);
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
						setState(624);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(20);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 5:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(627);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(628);
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
						setState(629);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(19);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 6:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(632);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(633);
						((ExprContext)_localctx).op = match(AND);
						setState(634);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(18);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 7:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(637);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(638);
						((ExprContext)_localctx).op = match(OR);
						setState(639);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(17);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					}
					} 
				}
				setState(646);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,39,_ctx);
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

	public static class IntfunctionstmtContext extends ParserRuleContext {
		public interfaces.Expression intfunc;
		public Token INTS;
		public ExprContext expr;
		public TerminalNode INTS() { return getToken(GrammarParser.INTS, 0); }
		public TerminalNode PARIZQ() { return getToken(GrammarParser.PARIZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(GrammarParser.PARDER, 0); }
		public IntfunctionstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_intfunctionstmt; }
	}

	public final IntfunctionstmtContext intfunctionstmt() throws RecognitionException {
		IntfunctionstmtContext _localctx = new IntfunctionstmtContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_intfunctionstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(647);
			((IntfunctionstmtContext)_localctx).INTS = match(INTS);
			setState(648);
			match(PARIZQ);
			setState(649);
			((IntfunctionstmtContext)_localctx).expr = expr(0);
			setState(650);
			match(PARDER);
			_localctx.intfunc = expressions.NewFunctionInt((((IntfunctionstmtContext)_localctx).INTS!=null?((IntfunctionstmtContext)_localctx).INTS.getLine():0), (((IntfunctionstmtContext)_localctx).INTS!=null?((IntfunctionstmtContext)_localctx).INTS.getCharPositionInLine():0), ((IntfunctionstmtContext)_localctx).expr.e)
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

	public static class FloatfunctionstmtContext extends ParserRuleContext {
		public interfaces.Expression floatfunc;
		public Token FLOATS;
		public ExprContext expr;
		public TerminalNode FLOATS() { return getToken(GrammarParser.FLOATS, 0); }
		public TerminalNode PARIZQ() { return getToken(GrammarParser.PARIZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(GrammarParser.PARDER, 0); }
		public FloatfunctionstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_floatfunctionstmt; }
	}

	public final FloatfunctionstmtContext floatfunctionstmt() throws RecognitionException {
		FloatfunctionstmtContext _localctx = new FloatfunctionstmtContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_floatfunctionstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(653);
			((FloatfunctionstmtContext)_localctx).FLOATS = match(FLOATS);
			setState(654);
			match(PARIZQ);
			setState(655);
			((FloatfunctionstmtContext)_localctx).expr = expr(0);
			setState(656);
			match(PARDER);
			_localctx.floatfunc = expressions.NewFunctionFloat((((FloatfunctionstmtContext)_localctx).FLOATS!=null?((FloatfunctionstmtContext)_localctx).FLOATS.getLine():0), (((FloatfunctionstmtContext)_localctx).FLOATS!=null?((FloatfunctionstmtContext)_localctx).FLOATS.getCharPositionInLine():0), ((FloatfunctionstmtContext)_localctx).expr.e)
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

	public static class StringfunctionstmtContext extends ParserRuleContext {
		public interfaces.Expression stringfunc;
		public Token STRINGS;
		public ExprContext expr;
		public TerminalNode STRINGS() { return getToken(GrammarParser.STRINGS, 0); }
		public TerminalNode PARIZQ() { return getToken(GrammarParser.PARIZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(GrammarParser.PARDER, 0); }
		public StringfunctionstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stringfunctionstmt; }
	}

	public final StringfunctionstmtContext stringfunctionstmt() throws RecognitionException {
		StringfunctionstmtContext _localctx = new StringfunctionstmtContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_stringfunctionstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(659);
			((StringfunctionstmtContext)_localctx).STRINGS = match(STRINGS);
			setState(660);
			match(PARIZQ);
			setState(661);
			((StringfunctionstmtContext)_localctx).expr = expr(0);
			setState(662);
			match(PARDER);
			_localctx.stringfunc = expressions.NewFunctionString((((StringfunctionstmtContext)_localctx).STRINGS!=null?((StringfunctionstmtContext)_localctx).STRINGS.getLine():0), (((StringfunctionstmtContext)_localctx).STRINGS!=null?((StringfunctionstmtContext)_localctx).STRINGS.getCharPositionInLine():0), ((StringfunctionstmtContext)_localctx).expr.e)
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

	public static class AccessfuncstmtContext extends ParserRuleContext {
		public interfaces.Expression access;
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode PARIZQ() { return getToken(GrammarParser.PARIZQ, 0); }
		public TerminalNode AND_SIMPLE() { return getToken(GrammarParser.AND_SIMPLE, 0); }
		public TerminalNode PARDER() { return getToken(GrammarParser.PARDER, 0); }
		public ListParamsContext listParams() {
			return getRuleContext(ListParamsContext.class,0);
		}
		public AccessfuncstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_accessfuncstmt; }
	}

	public final AccessfuncstmtContext accessfuncstmt() throws RecognitionException {
		AccessfuncstmtContext _localctx = new AccessfuncstmtContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_accessfuncstmt);
		try {
			setState(674);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,40,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(665);
				match(ID);
				setState(666);
				match(PARIZQ);
				setState(667);
				match(AND_SIMPLE);
				setState(668);
				match(PARDER);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(669);
				match(ID);
				setState(670);
				match(PARIZQ);
				setState(671);
				listParams(0);
				setState(672);
				match(PARDER);
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
		enterRule(_localctx, 68, RULE_accessstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(676);
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
		public TerminalNode PTCOMA() { return getToken(GrammarParser.PTCOMA, 0); }
		public TerminalNode IG_SUB() { return getToken(GrammarParser.IG_SUB, 0); }
		public IncreaseanddecreasestmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_increaseanddecreasestmt; }
	}

	public final IncreaseanddecreasestmtContext increaseanddecreasestmt() throws RecognitionException {
		IncreaseanddecreasestmtContext _localctx = new IncreaseanddecreasestmtContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_increaseanddecreasestmt);
		int _la;
		try {
			setState(695);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,43,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(679);
				((IncreaseanddecreasestmtContext)_localctx).ID = match(ID);
				setState(680);
				((IncreaseanddecreasestmtContext)_localctx).IG_ADD = match(IG_ADD);
				setState(681);
				((IncreaseanddecreasestmtContext)_localctx).expr = expr(0);
				setState(683);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(682);
					match(PTCOMA);
					}
				}

				_localctx.increasedecrease = instructions.NewIncreaseDecrease((((IncreaseanddecreasestmtContext)_localctx).ID!=null?((IncreaseanddecreasestmtContext)_localctx).ID.getLine():0), (((IncreaseanddecreasestmtContext)_localctx).ID!=null?((IncreaseanddecreasestmtContext)_localctx).ID.getCharPositionInLine():0), (((IncreaseanddecreasestmtContext)_localctx).ID!=null?((IncreaseanddecreasestmtContext)_localctx).ID.getText():null), (((IncreaseanddecreasestmtContext)_localctx).IG_ADD!=null?((IncreaseanddecreasestmtContext)_localctx).IG_ADD.getText():null), ((IncreaseanddecreasestmtContext)_localctx).expr.e)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(687);
				((IncreaseanddecreasestmtContext)_localctx).ID = match(ID);
				setState(688);
				((IncreaseanddecreasestmtContext)_localctx).IG_SUB = match(IG_SUB);
				setState(689);
				((IncreaseanddecreasestmtContext)_localctx).expr = expr(0);
				setState(691);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(690);
					match(PTCOMA);
					}
				}

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
		enterRule(_localctx, 72, RULE_type);
		try {
			setState(707);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRINGS:
				enterOuterAlt(_localctx, 1);
				{
				setState(697);
				match(STRINGS);
				_localctx.t = environment.STRING
				}
				break;
			case INTS:
				enterOuterAlt(_localctx, 2);
				{
				setState(699);
				match(INTS);
				_localctx.t = environment.INTEGER
				}
				break;
			case FLOATS:
				enterOuterAlt(_localctx, 3);
				{
				setState(701);
				match(FLOATS);
				_localctx.t = environment.FLOAT
				}
				break;
			case BOOLS:
				enterOuterAlt(_localctx, 4);
				{
				setState(703);
				match(BOOLS);
				_localctx.t = environment.BOOLEAN 
				}
				break;
			case CHARACTERS:
				enterOuterAlt(_localctx, 5);
				{
				setState(705);
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
		int _startState = 74;
		enterRecursionRule(_localctx, 74, RULE_listParams, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(710);
			((ListParamsContext)_localctx).expr = expr(0);

			            _localctx.l = []interface{}{}
			            _localctx.l = append(_localctx.l, ((ListParamsContext)_localctx).expr.e)
			        
			}
			_ctx.stop = _input.LT(-1);
			setState(720);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,45,_ctx);
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
					setState(713);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(714);
					match(COMA);
					setState(715);
					((ListParamsContext)_localctx).expr = expr(0);

					                                          var arr []interface{}
					                                          arr = append(((ListParamsContext)_localctx).list.l, ((ListParamsContext)_localctx).expr.e)
					                                          _localctx.l = arr
					                                      
					}
					} 
				}
				setState(722);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,45,_ctx);
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

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 27:
			return listParamsFunc_sempred((ListParamsFuncContext)_localctx, predIndex);
		case 29:
			return expr_sempred((ExprContext)_localctx, predIndex);
		case 37:
			return listParams_sempred((ListParamsContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean listParamsFunc_sempred(ListParamsFuncContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return precpred(_ctx, 22);
		case 2:
			return precpred(_ctx, 21);
		case 3:
			return precpred(_ctx, 20);
		case 4:
			return precpred(_ctx, 19);
		case 5:
			return precpred(_ctx, 18);
		case 6:
			return precpred(_ctx, 17);
		case 7:
			return precpred(_ctx, 16);
		}
		return true;
	}
	private boolean listParams_sempred(ListParamsContext _localctx, int predIndex) {
		switch (predIndex) {
		case 8:
			return precpred(_ctx, 2);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3E\u02d6\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\3\2\3\2\3\2\3\2\3\3\6\3T\n"+
		"\3\r\3\16\3U\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4\u0090\n\4\3\5\3\5\3\5\3\5\3\5\5\5\u0097"+
		"\n\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6\u00a2\n\6\3\6\3\6\3\6\3\6"+
		"\3\6\3\6\3\6\5\6\u00ab\n\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6\u00b5\n"+
		"\6\3\6\3\6\5\6\u00b9\n\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\5\7\u00c2\n\7\3\7"+
		"\3\7\3\7\3\7\3\7\3\7\3\7\5\7\u00cb\n\7\3\7\3\7\5\7\u00cf\n\7\3\b\6\b\u00d2"+
		"\n\b\r\b\16\b\u00d3\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\5\t"+
		"\u00f3\n\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n"+
		"\3\n\5\n\u0105\n\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3"+
		"\13\3\13\3\13\3\13\3\13\3\13\5\13\u0117\n\13\3\f\6\f\u011a\n\f\r\f\16"+
		"\f\u011b\3\f\3\f\3\r\3\r\3\r\3\r\5\r\u0124\n\r\3\r\3\r\3\16\3\16\3\16"+
		"\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17"+
		"\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\5\17\u0144\n\17"+
		"\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\21\3\21\5\21\u0151\n\21"+
		"\3\21\3\21\3\21\3\21\5\21\u0157\n\21\3\21\3\21\3\21\3\21\5\21\u015d\n"+
		"\21\3\21\3\21\3\21\5\21\u0162\n\21\3\21\5\21\u0165\n\21\3\22\3\22\3\22"+
		"\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\5\22\u0172\n\22\3\22\3\22\3\22"+
		"\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\5\22\u0180\n\22\3\22\3\22"+
		"\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\5\22\u018d\n\22\3\22\3\22"+
		"\3\22\3\22\3\22\3\22\3\22\3\22\3\22\5\22\u0198\n\22\3\22\3\22\5\22\u019c"+
		"\n\22\3\23\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\24\3\24\3\24\5\24"+
		"\u01aa\n\24\3\24\3\24\3\25\3\25\3\25\3\25\3\25\3\25\5\25\u01b4\n\25\3"+
		"\25\3\25\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\5\26\u01c1\n\26"+
		"\3\26\3\26\3\27\3\27\3\27\3\27\3\27\3\30\3\30\3\30\3\30\3\30\3\31\3\31"+
		"\3\31\3\31\3\31\3\31\3\32\3\32\3\32\3\32\3\32\3\32\3\32\5\32\u01dc\n\32"+
		"\3\32\3\32\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\5\33\u01e8\n\33\3\34"+
		"\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34"+
		"\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34"+
		"\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34"+
		"\3\34\5\34\u0216\n\34\3\35\3\35\3\35\3\35\3\35\3\35\3\35\7\35\u021f\n"+
		"\35\f\35\16\35\u0222\13\35\3\36\3\36\3\36\5\36\u0227\n\36\3\36\3\36\3"+
		"\36\3\36\3\36\5\36\u022e\n\36\3\36\5\36\u0231\n\36\3\37\3\37\3\37\3\37"+
		"\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37"+
		"\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37"+
		"\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\5\37"+
		"\u0260\n\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37"+
		"\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37"+
		"\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\7\37\u0285\n\37\f\37\16"+
		"\37\u0288\13\37\3 \3 \3 \3 \3 \3 \3!\3!\3!\3!\3!\3!\3\"\3\"\3\"\3\"\3"+
		"\"\3\"\3#\3#\3#\3#\3#\3#\3#\3#\3#\5#\u02a5\n#\3$\3$\3$\3%\3%\3%\3%\5%"+
		"\u02ae\n%\3%\3%\3%\3%\3%\3%\5%\u02b6\n%\3%\3%\5%\u02ba\n%\3&\3&\3&\3&"+
		"\3&\3&\3&\3&\3&\3&\5&\u02c6\n&\3\'\3\'\3\'\3\'\3\'\3\'\3\'\3\'\3\'\7\'"+
		"\u02d1\n\'\f\'\16\'\u02d4\13\'\3\'\2\58<L(\2\4\6\b\n\f\16\20\22\24\26"+
		"\30\32\34\36 \"$&(*,.\60\62\64\668:<>@BDFHJL\2\t\3\2\27\31\4\2$$AA\3\2"+
		"\61\62\3\2\63\65\4\2--//\4\2..\60\60\3\2%&\2\u030a\2N\3\2\2\2\4S\3\2\2"+
		"\2\6\u008f\3\2\2\2\b\u0091\3\2\2\2\n\u00b8\3\2\2\2\f\u00ce\3\2\2\2\16"+
		"\u00d1\3\2\2\2\20\u00f2\3\2\2\2\22\u0104\3\2\2\2\24\u0116\3\2\2\2\26\u0119"+
		"\3\2\2\2\30\u011f\3\2\2\2\32\u0127\3\2\2\2\34\u0143\3\2\2\2\36\u0145\3"+
		"\2\2\2 \u0164\3\2\2\2\"\u019b\3\2\2\2$\u019d\3\2\2\2&\u01a2\3\2\2\2(\u01ad"+
		"\3\2\2\2*\u01b7\3\2\2\2,\u01c4\3\2\2\2.\u01c9\3\2\2\2\60\u01ce\3\2\2\2"+
		"\62\u01d4\3\2\2\2\64\u01df\3\2\2\2\66\u0215\3\2\2\28\u0217\3\2\2\2:\u0230"+
		"\3\2\2\2<\u025f\3\2\2\2>\u0289\3\2\2\2@\u028f\3\2\2\2B\u0295\3\2\2\2D"+
		"\u02a4\3\2\2\2F\u02a6\3\2\2\2H\u02b9\3\2\2\2J\u02c5\3\2\2\2L\u02c7\3\2"+
		"\2\2NO\5\4\3\2OP\7\2\2\3PQ\b\2\1\2Q\3\3\2\2\2RT\5\6\4\2SR\3\2\2\2TU\3"+
		"\2\2\2US\3\2\2\2UV\3\2\2\2VW\3\2\2\2WX\b\3\1\2X\5\3\2\2\2YZ\5\b\5\2Z["+
		"\b\4\1\2[\u0090\3\2\2\2\\]\5\n\6\2]^\b\4\1\2^\u0090\3\2\2\2_`\5\f\7\2"+
		"`a\b\4\1\2a\u0090\3\2\2\2bc\5\30\r\2cd\b\4\1\2d\u0090\3\2\2\2ef\5H%\2"+
		"fg\b\4\1\2g\u0090\3\2\2\2hi\5\20\t\2ij\b\4\1\2j\u0090\3\2\2\2kl\5\32\16"+
		"\2lm\b\4\1\2m\u0090\3\2\2\2no\5\34\17\2op\b\4\1\2p\u0090\3\2\2\2qr\5\22"+
		"\n\2rs\b\4\1\2s\u0090\3\2\2\2tu\5\36\20\2uv\b\4\1\2v\u0090\3\2\2\2wx\5"+
		" \21\2xy\b\4\1\2y\u0090\3\2\2\2z{\5\"\22\2{|\b\4\1\2|\u0090\3\2\2\2}~"+
		"\5&\24\2~\177\b\4\1\2\177\u0090\3\2\2\2\u0080\u0081\5(\25\2\u0081\u0082"+
		"\b\4\1\2\u0082\u0090\3\2\2\2\u0083\u0084\5*\26\2\u0084\u0085\b\4\1\2\u0085"+
		"\u0090\3\2\2\2\u0086\u0087\5\62\32\2\u0087\u0088\b\4\1\2\u0088\u0090\3"+
		"\2\2\2\u0089\u008a\5\66\34\2\u008a\u008b\b\4\1\2\u008b\u0090\3\2\2\2\u008c"+
		"\u008d\5$\23\2\u008d\u008e\b\4\1\2\u008e\u0090\3\2\2\2\u008fY\3\2\2\2"+
		"\u008f\\\3\2\2\2\u008f_\3\2\2\2\u008fb\3\2\2\2\u008fe\3\2\2\2\u008fh\3"+
		"\2\2\2\u008fk\3\2\2\2\u008fn\3\2\2\2\u008fq\3\2\2\2\u008ft\3\2\2\2\u008f"+
		"w\3\2\2\2\u008fz\3\2\2\2\u008f}\3\2\2\2\u008f\u0080\3\2\2\2\u008f\u0083"+
		"\3\2\2\2\u008f\u0086\3\2\2\2\u008f\u0089\3\2\2\2\u008f\u008c\3\2\2\2\u0090"+
		"\7\3\2\2\2\u0091\u0092\7\f\2\2\u0092\u0093\7\66\2\2\u0093\u0094\5L\'\2"+
		"\u0094\u0096\7\67\2\2\u0095\u0097\7>\2\2\u0096\u0095\3\2\2\2\u0096\u0097"+
		"\3\2\2\2\u0097\u0098\3\2\2\2\u0098\u0099\b\5\1\2\u0099\t\3\2\2\2\u009a"+
		"\u009b\7\b\2\2\u009b\u009c\7$\2\2\u009c\u009d\7:\2\2\u009d\u009e\5J&\2"+
		"\u009e\u009f\7*\2\2\u009f\u00a1\5<\37\2\u00a0\u00a2\7>\2\2\u00a1\u00a0"+
		"\3\2\2\2\u00a1\u00a2\3\2\2\2\u00a2\u00a3\3\2\2\2\u00a3\u00a4\b\6\1\2\u00a4"+
		"\u00b9\3\2\2\2\u00a5\u00a6\7\b\2\2\u00a6\u00a7\7$\2\2\u00a7\u00a8\7*\2"+
		"\2\u00a8\u00aa\5<\37\2\u00a9\u00ab\7>\2\2\u00aa\u00a9\3\2\2\2\u00aa\u00ab"+
		"\3\2\2\2\u00ab\u00ac\3\2\2\2\u00ac\u00ad\b\6\1\2\u00ad\u00b9\3\2\2\2\u00ae"+
		"\u00af\7\b\2\2\u00af\u00b0\7$\2\2\u00b0\u00b1\7:\2\2\u00b1\u00b2\5J&\2"+
		"\u00b2\u00b4\7;\2\2\u00b3\u00b5\7>\2\2\u00b4\u00b3\3\2\2\2\u00b4\u00b5"+
		"\3\2\2\2\u00b5\u00b6\3\2\2\2\u00b6\u00b7\b\6\1\2\u00b7\u00b9\3\2\2\2\u00b8"+
		"\u009a\3\2\2\2\u00b8\u00a5\3\2\2\2\u00b8\u00ae\3\2\2\2\u00b9\13\3\2\2"+
		"\2\u00ba\u00bb\7\20\2\2\u00bb\u00bc\7$\2\2\u00bc\u00bd\7:\2\2\u00bd\u00be"+
		"\5J&\2\u00be\u00bf\7*\2\2\u00bf\u00c1\5<\37\2\u00c0\u00c2\7>\2\2\u00c1"+
		"\u00c0\3\2\2\2\u00c1\u00c2\3\2\2\2\u00c2\u00c3\3\2\2\2\u00c3\u00c4\b\7"+
		"\1\2\u00c4\u00cf\3\2\2\2\u00c5\u00c6\7\20\2\2\u00c6\u00c7\7$\2\2\u00c7"+
		"\u00c8\7*\2\2\u00c8\u00ca\5<\37\2\u00c9\u00cb\7>\2\2\u00ca\u00c9\3\2\2"+
		"\2\u00ca\u00cb\3\2\2\2\u00cb\u00cc\3\2\2\2\u00cc\u00cd\b\7\1\2\u00cd\u00cf"+
		"\3\2\2\2\u00ce\u00ba\3\2\2\2\u00ce\u00c5\3\2\2\2\u00cf\r\3\2\2\2\u00d0"+
		"\u00d2\5\20\t\2\u00d1\u00d0\3\2\2\2\u00d2\u00d3\3\2\2\2\u00d3\u00d1\3"+
		"\2\2\2\u00d3\u00d4\3\2\2\2\u00d4\u00d5\3\2\2\2\u00d5\u00d6\b\b\1\2\u00d6"+
		"\17\3\2\2\2\u00d7\u00d8\7\r\2\2\u00d8\u00d9\5<\37\2\u00d9\u00da\78\2\2"+
		"\u00da\u00db\5\4\3\2\u00db\u00dc\79\2\2\u00dc\u00dd\b\t\1\2\u00dd\u00f3"+
		"\3\2\2\2\u00de\u00df\7\r\2\2\u00df\u00e0\5<\37\2\u00e0\u00e1\78\2\2\u00e1"+
		"\u00e2\5\4\3\2\u00e2\u00e3\79\2\2\u00e3\u00e4\7\16\2\2\u00e4\u00e5\78"+
		"\2\2\u00e5\u00e6\5\4\3\2\u00e6\u00e7\79\2\2\u00e7\u00e8\b\t\1\2\u00e8"+
		"\u00f3\3\2\2\2\u00e9\u00ea\7\r\2\2\u00ea\u00eb\5<\37\2\u00eb\u00ec\78"+
		"\2\2\u00ec\u00ed\5\4\3\2\u00ed\u00ee\79\2\2\u00ee\u00ef\7\16\2\2\u00ef"+
		"\u00f0\5\16\b\2\u00f0\u00f1\b\t\1\2\u00f1\u00f3\3\2\2\2\u00f2\u00d7\3"+
		"\2\2\2\u00f2\u00de\3\2\2\2\u00f2\u00e9\3\2\2\2\u00f3\21\3\2\2\2\u00f4"+
		"\u00f5\7\23\2\2\u00f5\u00f6\5<\37\2\u00f6\u00f7\78\2\2\u00f7\u00f8\7\25"+
		"\2\2\u00f8\u00f9\7:\2\2\u00f9\u00fa\5\4\3\2\u00fa\u00fb\79\2\2\u00fb\u00fc"+
		"\b\n\1\2\u00fc\u0105\3\2\2\2\u00fd\u00fe\7\23\2\2\u00fe\u00ff\5<\37\2"+
		"\u00ff\u0100\78\2\2\u0100\u0101\5\24\13\2\u0101\u0102\79\2\2\u0102\u0103"+
		"\b\n\1\2\u0103\u0105\3\2\2\2\u0104\u00f4\3\2\2\2\u0104\u00fd\3\2\2\2\u0105"+
		"\23\3\2\2\2\u0106\u0107\7\24\2\2\u0107\u0108\5<\37\2\u0108\u0109\7:\2"+
		"\2\u0109\u010a\5\4\3\2\u010a\u010b\5\26\f\2\u010b\u010c\b\13\1\2\u010c"+
		"\u0117\3\2\2\2\u010d\u010e\7\24\2\2\u010e\u010f\5<\37\2\u010f\u0110\7"+
		":\2\2\u0110\u0111\5\4\3\2\u0111\u0112\7\25\2\2\u0112\u0113\7:\2\2\u0113"+
		"\u0114\5\4\3\2\u0114\u0115\b\13\1\2\u0115\u0117\3\2\2\2\u0116\u0106\3"+
		"\2\2\2\u0116\u010d\3\2\2\2\u0117\25\3\2\2\2\u0118\u011a\5\24\13\2\u0119"+
		"\u0118\3\2\2\2\u011a\u011b\3\2\2\2\u011b\u0119\3\2\2\2\u011b\u011c\3\2"+
		"\2\2\u011c\u011d\3\2\2\2\u011d\u011e\b\f\1\2\u011e\27\3\2\2\2\u011f\u0120"+
		"\7$\2\2\u0120\u0121\7*\2\2\u0121\u0123\5<\37\2\u0122\u0124\7>\2\2\u0123"+
		"\u0122\3\2\2\2\u0123\u0124\3\2\2\2\u0124\u0125\3\2\2\2\u0125\u0126\b\r"+
		"\1\2\u0126\31\3\2\2\2\u0127\u0128\7\17\2\2\u0128\u0129\5<\37\2\u0129\u012a"+
		"\78\2\2\u012a\u012b\5\4\3\2\u012b\u012c\79\2\2\u012c\u012d\b\16\1\2\u012d"+
		"\33\3\2\2\2\u012e\u012f\7\21\2\2\u012f\u0130\7$\2\2\u0130\u0131\7\22\2"+
		"\2\u0131\u0132\5<\37\2\u0132\u0133\78\2\2\u0133\u0134\5\4\3\2\u0134\u0135"+
		"\79\2\2\u0135\u0136\b\17\1\2\u0136\u0144\3\2\2\2\u0137\u0138\7\21\2\2"+
		"\u0138\u0139\7$\2\2\u0139\u013a\7\22\2\2\u013a\u013b\5<\37\2\u013b\u013c"+
		"\7<\2\2\u013c\u013d\7<\2\2\u013d\u013e\5<\37\2\u013e\u013f\78\2\2\u013f"+
		"\u0140\5\4\3\2\u0140\u0141\79\2\2\u0141\u0142\b\17\1\2\u0142\u0144\3\2"+
		"\2\2\u0143\u012e\3\2\2\2\u0143\u0137\3\2\2\2\u0144\35\3\2\2\2\u0145\u0146"+
		"\7\26\2\2\u0146\u0147\5<\37\2\u0147\u0148\7\16\2\2\u0148\u0149\78\2\2"+
		"\u0149\u014a\5\4\3\2\u014a\u014b\t\2\2\2\u014b\u014c\79\2\2\u014c\u014d"+
		"\b\20\1\2\u014d\37\3\2\2\2\u014e\u0150\7\30\2\2\u014f\u0151\7>\2\2\u0150"+
		"\u014f\3\2\2\2\u0150\u0151\3\2\2\2\u0151\u0152\3\2\2\2\u0152\u0165\b\21"+
		"\1\2\u0153\u0154\7\30\2\2\u0154\u0156\5<\37\2\u0155\u0157\7>\2\2\u0156"+
		"\u0155\3\2\2\2\u0156\u0157\3\2\2\2\u0157\u0158\3\2\2\2\u0158\u0159\b\21"+
		"\1\2\u0159\u0165\3\2\2\2\u015a\u015c\7\27\2\2\u015b\u015d\7>\2\2\u015c"+
		"\u015b\3\2\2\2\u015c\u015d\3\2\2\2\u015d\u015e\3\2\2\2\u015e\u0165\b\21"+
		"\1\2\u015f\u0161\7\31\2\2\u0160\u0162\7>\2\2\u0161\u0160\3\2\2\2\u0161"+
		"\u0162\3\2\2\2\u0162\u0163\3\2\2\2\u0163\u0165\b\21\1\2\u0164\u014e\3"+
		"\2\2\2\u0164\u0153\3\2\2\2\u0164\u015a\3\2\2\2\u0164\u015f\3\2\2\2\u0165"+
		"!\3\2\2\2\u0166\u0167\7\b\2\2\u0167\u0168\7$\2\2\u0168\u0169\7:\2\2\u0169"+
		"\u016a\7?\2\2\u016a\u016b\5J&\2\u016b\u016c\7@\2\2\u016c\u016d\7*\2\2"+
		"\u016d\u016e\7?\2\2\u016e\u016f\5L\'\2\u016f\u0171\7@\2\2\u0170\u0172"+
		"\7>\2\2\u0171\u0170\3\2\2\2\u0171\u0172\3\2\2\2\u0172\u0173\3\2\2\2\u0173"+
		"\u0174\b\22\1\2\u0174\u019c\3\2\2\2\u0175\u0176\7\b\2\2\u0176\u0177\7"+
		"$\2\2\u0177\u0178\7:\2\2\u0178\u0179\7?\2\2\u0179\u017a\5J&\2\u017a\u017b"+
		"\7@\2\2\u017b\u017c\7*\2\2\u017c\u017d\7?\2\2\u017d\u017f\7@\2\2\u017e"+
		"\u0180\7>\2\2\u017f\u017e\3\2\2\2\u017f\u0180\3\2\2\2\u0180\u0181\3\2"+
		"\2\2\u0181\u0182\b\22\1\2\u0182\u019c\3\2\2\2\u0183\u0184\7\b\2\2\u0184"+
		"\u0185\7$\2\2\u0185\u0186\7*\2\2\u0186\u0187\7?\2\2\u0187\u0188\5J&\2"+
		"\u0188\u0189\7@\2\2\u0189\u018a\7\66\2\2\u018a\u018c\7\67\2\2\u018b\u018d"+
		"\7>\2\2\u018c\u018b\3\2\2\2\u018c\u018d\3\2\2\2\u018d\u019c\3\2\2\2\u018e"+
		"\u018f\7\b\2\2\u018f\u0190\7$\2\2\u0190\u0191\7:\2\2\u0191\u0192\7?\2"+
		"\2\u0192\u0193\5J&\2\u0193\u0194\7@\2\2\u0194\u0195\7*\2\2\u0195\u0197"+
		"\7$\2\2\u0196\u0198\7>\2\2\u0197\u0196\3\2\2\2\u0197\u0198\3\2\2\2\u0198"+
		"\u0199\3\2\2\2\u0199\u019a\b\22\1\2\u019a\u019c\3\2\2\2\u019b\u0166\3"+
		"\2\2\2\u019b\u0175\3\2\2\2\u019b\u0183\3\2\2\2\u019b\u018e\3\2\2\2\u019c"+
		"#\3\2\2\2\u019d\u019e\7$\2\2\u019e\u019f\7\66\2\2\u019f\u01a0\7\67\2\2"+
		"\u01a0\u01a1\b\23\1\2\u01a1%\3\2\2\2\u01a2\u01a3\7$\2\2\u01a3\u01a4\7"+
		"<\2\2\u01a4\u01a5\7\32\2\2\u01a5\u01a6\7\66\2\2\u01a6\u01a7\5<\37\2\u01a7"+
		"\u01a9\7\67\2\2\u01a8\u01aa\7>\2\2\u01a9\u01a8\3\2\2\2\u01a9\u01aa\3\2"+
		"\2\2\u01aa\u01ab\3\2\2\2\u01ab\u01ac\b\24\1\2\u01ac\'\3\2\2\2\u01ad\u01ae"+
		"\7$\2\2\u01ae\u01af\7<\2\2\u01af\u01b0\7\33\2\2\u01b0\u01b1\7\66\2\2\u01b1"+
		"\u01b3\7\67\2\2\u01b2\u01b4\7>\2\2\u01b3\u01b2\3\2\2\2\u01b3\u01b4\3\2"+
		"\2\2\u01b4\u01b5\3\2\2\2\u01b5\u01b6\b\25\1\2\u01b6)\3\2\2\2\u01b7\u01b8"+
		"\7$\2\2\u01b8\u01b9\7<\2\2\u01b9\u01ba\7\34\2\2\u01ba\u01bb\7\66\2\2\u01bb"+
		"\u01bc\7\35\2\2\u01bc\u01bd\7:\2\2\u01bd\u01be\5<\37\2\u01be\u01c0\7\67"+
		"\2\2\u01bf\u01c1\7>\2\2\u01c0\u01bf\3\2\2\2\u01c0\u01c1\3\2\2\2\u01c1"+
		"\u01c2\3\2\2\2\u01c2\u01c3\b\26\1\2\u01c3+\3\2\2\2\u01c4\u01c5\7$\2\2"+
		"\u01c5\u01c6\7<\2\2\u01c6\u01c7\7\36\2\2\u01c7\u01c8\b\27\1\2\u01c8-\3"+
		"\2\2\2\u01c9\u01ca\7$\2\2\u01ca\u01cb\7<\2\2\u01cb\u01cc\7\37\2\2\u01cc"+
		"\u01cd\b\30\1\2\u01cd/\3\2\2\2\u01ce\u01cf\7$\2\2\u01cf\u01d0\7?\2\2\u01d0"+
		"\u01d1\5<\37\2\u01d1\u01d2\7@\2\2\u01d2\u01d3\b\31\1\2\u01d3\61\3\2\2"+
		"\2\u01d4\u01d5\7$\2\2\u01d5\u01d6\7?\2\2\u01d6\u01d7\5<\37\2\u01d7\u01d8"+
		"\7@\2\2\u01d8\u01d9\7*\2\2\u01d9\u01db\5<\37\2\u01da\u01dc\7>\2\2\u01db"+
		"\u01da\3\2\2\2\u01db\u01dc\3\2\2\2\u01dc\u01dd\3\2\2\2\u01dd\u01de\b\32"+
		"\1\2\u01de\63\3\2\2\2\u01df\u01e0\7\b\2\2\u01e0\u01e1\7$\2\2\u01e1\u01e2"+
		"\7\66\2\2\u01e2\u01e3\7:\2\2\u01e3\u01e4\5J&\2\u01e4\u01e5\7\67\2\2\u01e5"+
		"\u01e7\7*\2\2\u01e6\u01e8\7>\2\2\u01e7\u01e6\3\2\2\2\u01e7\u01e8\3\2\2"+
		"\2\u01e8\65\3\2\2\2\u01e9\u01ea\7 \2\2\u01ea\u01eb\7$\2\2\u01eb\u01ec"+
		"\7\66\2\2\u01ec\u01ed\5L\'\2\u01ed\u01ee\7\67\2\2\u01ee\u01ef\7\64\2\2"+
		"\u01ef\u01f0\7/\2\2\u01f0\u01f1\5J&\2\u01f1\u01f2\78\2\2\u01f2\u01f3\5"+
		"\4\3\2\u01f3\u01f4\79\2\2\u01f4\u01f5\b\34\1\2\u01f5\u0216\3\2\2\2\u01f6"+
		"\u01f7\7 \2\2\u01f7\u01f8\7$\2\2\u01f8\u01f9\7\66\2\2\u01f9\u01fa\5L\'"+
		"\2\u01fa\u01fb\7\67\2\2\u01fb\u01fc\78\2\2\u01fc\u01fd\5\4\3\2\u01fd\u01fe"+
		"\79\2\2\u01fe\u01ff\b\34\1\2\u01ff\u0216\3\2\2\2\u0200\u0201\7 \2\2\u0201"+
		"\u0202\7$\2\2\u0202\u0203\7\66\2\2\u0203\u0204\7\67\2\2\u0204\u0205\7"+
		"\64\2\2\u0205\u0206\7/\2\2\u0206\u0207\5J&\2\u0207\u0208\78\2\2\u0208"+
		"\u0209\5\4\3\2\u0209\u020a\79\2\2\u020a\u020b\b\34\1\2\u020b\u0216\3\2"+
		"\2\2\u020c\u020d\7 \2\2\u020d\u020e\7$\2\2\u020e\u020f\7\66\2\2\u020f"+
		"\u0210\7\67\2\2\u0210\u0211\78\2\2\u0211\u0212\5\4\3\2\u0212\u0213\79"+
		"\2\2\u0213\u0214\b\34\1\2\u0214\u0216\3\2\2\2\u0215\u01e9\3\2\2\2\u0215"+
		"\u01f6\3\2\2\2\u0215\u0200\3\2\2\2\u0215\u020c\3\2\2\2\u0216\67\3\2\2"+
		"\2\u0217\u0218\b\35\1\2\u0218\u0219\5:\36\2\u0219\u021a\b\35\1\2\u021a"+
		"\u0220\3\2\2\2\u021b\u021c\f\4\2\2\u021c\u021d\7=\2\2\u021d\u021f\b\35"+
		"\1\2\u021e\u021b\3\2\2\2\u021f\u0222\3\2\2\2\u0220\u021e\3\2\2\2\u0220"+
		"\u0221\3\2\2\2\u02219\3\2\2\2\u0222\u0220\3\2\2\2\u0223\u0224\7$\2\2\u0224"+
		"\u0226\7:\2\2\u0225\u0227\7!\2\2\u0226\u0225\3\2\2\2\u0226\u0227\3\2\2"+
		"\2\u0227\u0228\3\2\2\2\u0228\u0231\5J&\2\u0229\u022a\t\3\2\2\u022a\u022b"+
		"\7$\2\2\u022b\u022d\7:\2\2\u022c\u022e\7!\2\2\u022d\u022c\3\2\2\2\u022d"+
		"\u022e\3\2\2\2\u022e\u022f\3\2\2\2\u022f\u0231\5J&\2\u0230\u0223\3\2\2"+
		"\2\u0230\u0229\3\2\2\2\u0231;\3\2\2\2\u0232\u0233\b\37\1\2\u0233\u0234"+
		"\7\'\2\2\u0234\u0235\5<\37\21\u0235\u0236\b\37\1\2\u0236\u0260\3\2\2\2"+
		"\u0237\u0238\7\64\2\2\u0238\u0239\5<\37\20\u0239\u023a\b\37\1\2\u023a"+
		"\u0260\3\2\2\2\u023b\u023c\7\66\2\2\u023c\u023d\5<\37\2\u023d\u023e\7"+
		"\67\2\2\u023e\u023f\b\37\1\2\u023f\u0260\3\2\2\2\u0240\u0241\7\"\2\2\u0241"+
		"\u0260\b\37\1\2\u0242\u0243\7#\2\2\u0243\u0260\b\37\1\2\u0244\u0245\7"+
		"\n\2\2\u0245\u0260\b\37\1\2\u0246\u0247\7\13\2\2\u0247\u0260\b\37\1\2"+
		"\u0248\u0249\7\t\2\2\u0249\u0260\b\37\1\2\u024a\u024b\5F$\2\u024b\u024c"+
		"\b\37\1\2\u024c\u0260\3\2\2\2\u024d\u024e\5,\27\2\u024e\u024f\b\37\1\2"+
		"\u024f\u0260\3\2\2\2\u0250\u0251\5.\30\2\u0251\u0252\b\37\1\2\u0252\u0260"+
		"\3\2\2\2\u0253\u0254\5\60\31\2\u0254\u0255\b\37\1\2\u0255\u0260\3\2\2"+
		"\2\u0256\u0257\5> \2\u0257\u0258\b\37\1\2\u0258\u0260\3\2\2\2\u0259\u025a"+
		"\5@!\2\u025a\u025b\b\37\1\2\u025b\u0260\3\2\2\2\u025c\u025d\5B\"\2\u025d"+
		"\u025e\b\37\1\2\u025e\u0260\3\2\2\2\u025f\u0232\3\2\2\2\u025f\u0237\3"+
		"\2\2\2\u025f\u023b\3\2\2\2\u025f\u0240\3\2\2\2\u025f\u0242\3\2\2\2\u025f"+
		"\u0244\3\2\2\2\u025f\u0246\3\2\2\2\u025f\u0248\3\2\2\2\u025f\u024a\3\2"+
		"\2\2\u025f\u024d\3\2\2\2\u025f\u0250\3\2\2\2\u025f\u0253\3\2\2\2\u025f"+
		"\u0256\3\2\2\2\u025f\u0259\3\2\2\2\u025f\u025c\3\2\2\2\u0260\u0286\3\2"+
		"\2\2\u0261\u0262\f\30\2\2\u0262\u0263\t\4\2\2\u0263\u0264\5<\37\31\u0264"+
		"\u0265\b\37\1\2\u0265\u0285\3\2\2\2\u0266\u0267\f\27\2\2\u0267\u0268\t"+
		"\5\2\2\u0268\u0269\5<\37\30\u0269\u026a\b\37\1\2\u026a\u0285\3\2\2\2\u026b"+
		"\u026c\f\26\2\2\u026c\u026d\t\6\2\2\u026d\u026e\5<\37\27\u026e\u026f\b"+
		"\37\1\2\u026f\u0285\3\2\2\2\u0270\u0271\f\25\2\2\u0271\u0272\t\7\2\2\u0272"+
		"\u0273\5<\37\26\u0273\u0274\b\37\1\2\u0274\u0285\3\2\2\2\u0275\u0276\f"+
		"\24\2\2\u0276\u0277\t\b\2\2\u0277\u0278\5<\37\25\u0278\u0279\b\37\1\2"+
		"\u0279\u0285\3\2\2\2\u027a\u027b\f\23\2\2\u027b\u027c\7)\2\2\u027c\u027d"+
		"\5<\37\24\u027d\u027e\b\37\1\2\u027e\u0285\3\2\2\2\u027f\u0280\f\22\2"+
		"\2\u0280\u0281\7(\2\2\u0281\u0282\5<\37\23\u0282\u0283\b\37\1\2\u0283"+
		"\u0285\3\2\2\2\u0284\u0261\3\2\2\2\u0284\u0266\3\2\2\2\u0284\u026b\3\2"+
		"\2\2\u0284\u0270\3\2\2\2\u0284\u0275\3\2\2\2\u0284\u027a\3\2\2\2\u0284"+
		"\u027f\3\2\2\2\u0285\u0288\3\2\2\2\u0286\u0284\3\2\2\2\u0286\u0287\3\2"+
		"\2\2\u0287=\3\2\2\2\u0288\u0286\3\2\2\2\u0289\u028a\7\4\2\2\u028a\u028b"+
		"\7\66\2\2\u028b\u028c\5<\37\2\u028c\u028d\7\67\2\2\u028d\u028e\b \1\2"+
		"\u028e?\3\2\2\2\u028f\u0290\7\5\2\2\u0290\u0291\7\66\2\2\u0291\u0292\5"+
		"<\37\2\u0292\u0293\7\67\2\2\u0293\u0294\b!\1\2\u0294A\3\2\2\2\u0295\u0296"+
		"\7\3\2\2\u0296\u0297\7\66\2\2\u0297\u0298\5<\37\2\u0298\u0299\7\67\2\2"+
		"\u0299\u029a\b\"\1\2\u029aC\3\2\2\2\u029b\u029c\7$\2\2\u029c\u029d\7\66"+
		"\2\2\u029d\u029e\7B\2\2\u029e\u02a5\7\67\2\2\u029f\u02a0\7$\2\2\u02a0"+
		"\u02a1\7\66\2\2\u02a1\u02a2\5L\'\2\u02a2\u02a3\7\67\2\2\u02a3\u02a5\3"+
		"\2\2\2\u02a4\u029b\3\2\2\2\u02a4\u029f\3\2\2\2\u02a5E\3\2\2\2\u02a6\u02a7"+
		"\7$\2\2\u02a7\u02a8\b$\1\2\u02a8G\3\2\2\2\u02a9\u02aa\7$\2\2\u02aa\u02ab"+
		"\7+\2\2\u02ab\u02ad\5<\37\2\u02ac\u02ae\7>\2\2\u02ad\u02ac\3\2\2\2\u02ad"+
		"\u02ae\3\2\2\2\u02ae\u02af\3\2\2\2\u02af\u02b0\b%\1\2\u02b0\u02ba\3\2"+
		"\2\2\u02b1\u02b2\7$\2\2\u02b2\u02b3\7,\2\2\u02b3\u02b5\5<\37\2\u02b4\u02b6"+
		"\7>\2\2\u02b5\u02b4\3\2\2\2\u02b5\u02b6\3\2\2\2\u02b6\u02b7\3\2\2\2\u02b7"+
		"\u02b8\b%\1\2\u02b8\u02ba\3\2\2\2\u02b9\u02a9\3\2\2\2\u02b9\u02b1\3\2"+
		"\2\2\u02baI\3\2\2\2\u02bb\u02bc\7\3\2\2\u02bc\u02c6\b&\1\2\u02bd\u02be"+
		"\7\4\2\2\u02be\u02c6\b&\1\2\u02bf\u02c0\7\5\2\2\u02c0\u02c6\b&\1\2\u02c1"+
		"\u02c2\7\6\2\2\u02c2\u02c6\b&\1\2\u02c3\u02c4\7\7\2\2\u02c4\u02c6\b&\1"+
		"\2\u02c5\u02bb\3\2\2\2\u02c5\u02bd\3\2\2\2\u02c5\u02bf\3\2\2\2\u02c5\u02c1"+
		"\3\2\2\2\u02c5\u02c3\3\2\2\2\u02c6K\3\2\2\2\u02c7\u02c8\b\'\1\2\u02c8"+
		"\u02c9\5<\37\2\u02c9\u02ca\b\'\1\2\u02ca\u02d2\3\2\2\2\u02cb\u02cc\f\4"+
		"\2\2\u02cc\u02cd\7=\2\2\u02cd\u02ce\5<\37\2\u02ce\u02cf\b\'\1\2\u02cf"+
		"\u02d1\3\2\2\2\u02d0\u02cb\3\2\2\2\u02d1\u02d4\3\2\2\2\u02d2\u02d0\3\2"+
		"\2\2\u02d2\u02d3\3\2\2\2\u02d3M\3\2\2\2\u02d4\u02d2\3\2\2\2\60U\u008f"+
		"\u0096\u00a1\u00aa\u00b4\u00b8\u00c1\u00ca\u00ce\u00d3\u00f2\u0104\u0116"+
		"\u011b\u0123\u0143\u0150\u0156\u015c\u0161\u0164\u0171\u017f\u018c\u0197"+
		"\u019b\u01a9\u01b3\u01c0\u01db\u01e7\u0215\u0220\u0226\u022d\u0230\u025f"+
		"\u0284\u0286\u02a4\u02ad\u02b5\u02b9\u02c5\u02d2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
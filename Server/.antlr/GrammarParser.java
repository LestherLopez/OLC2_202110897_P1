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
		STRUCT=32, NUMBER=33, STRING=34, ID=35, DIF=36, IG_IG=37, NOT=38, OR=39, 
		AND=40, IG=41, IG_ADD=42, IG_SUB=43, MAY_IG=44, MEN_IG=45, MAYOR=46, MENOR=47, 
		MUL=48, DIV=49, ADD=50, SUB=51, PERCENT=52, PARIZQ=53, PARDER=54, LLAVEIZQ=55, 
		LLAVEDER=56, DOUBLEPTS=57, QUESTION=58, POINT=59, COMA=60, PTCOMA=61, 
		CORCHETEIZQ=62, CORCHETEDER=63, GUION_BAJO=64, AND_SIMPLE=65, WHITESPACE=66, 
		COMMENT=67, LINE_COMMENT=68;
	public static final int
		RULE_s = 0, RULE_block = 1, RULE_instruction = 2, RULE_printstmt = 3, 
		RULE_declarestmt = 4, RULE_constantstmt = 5, RULE_blockelifs = 6, RULE_ifstmt = 7, 
		RULE_switchstmt = 8, RULE_casestmt = 9, RULE_blockcases = 10, RULE_assignationstmt = 11, 
		RULE_whilestmt = 12, RULE_forstmt = 13, RULE_guardstmt = 14, RULE_transferstmt = 15, 
		RULE_declarevectorstmt = 16, RULE_accessfuncinstruction = 17, RULE_appendstmt = 18, 
		RULE_removelaststmt = 19, RULE_removestmt = 20, RULE_emptvecstmt = 21, 
		RULE_countvecstmt = 22, RULE_accessvecstmt = 23, RULE_assignationvecstmt = 24, 
		RULE_declarematrixstmt = 25, RULE_declarefuncstmt = 26, RULE_listParamsFunc = 27, 
		RULE_parameterfuncstmt = 28, RULE_declarestructstmt = 29, RULE_listStruct = 30, 
		RULE_expr = 31, RULE_accessstructstmt = 32, RULE_accessfuncstmt = 33, 
		RULE_intfunctionstmt = 34, RULE_floatfunctionstmt = 35, RULE_stringfunctionstmt = 36, 
		RULE_accessstmt = 37, RULE_increaseanddecreasestmt = 38, RULE_type = 39, 
		RULE_listParams = 40, RULE_listStructExp = 41;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "block", "instruction", "printstmt", "declarestmt", "constantstmt", 
			"blockelifs", "ifstmt", "switchstmt", "casestmt", "blockcases", "assignationstmt", 
			"whilestmt", "forstmt", "guardstmt", "transferstmt", "declarevectorstmt", 
			"accessfuncinstruction", "appendstmt", "removelaststmt", "removestmt", 
			"emptvecstmt", "countvecstmt", "accessvecstmt", "assignationvecstmt", 
			"declarematrixstmt", "declarefuncstmt", "listParamsFunc", "parameterfuncstmt", 
			"declarestructstmt", "listStruct", "expr", "accessstructstmt", "accessfuncstmt", 
			"intfunctionstmt", "floatfunctionstmt", "stringfunctionstmt", "accessstmt", 
			"increaseanddecreasestmt", "type", "listParams", "listStructExp"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'String'", "'Int'", "'Float'", "'Bool'", "'Character'", "'var'", 
			"'nil'", "'true'", "'false'", "'print'", "'if'", "'else'", "'while'", 
			"'let'", "'for'", "'in'", "'switch'", "'case'", "'default'", "'guard'", 
			"'continue'", "'return'", "'break'", "'append'", "'removeLast'", "'remove'", 
			"'at'", "'IsEmpty'", "'count'", "'func'", "'inout'", "'struct'", null, 
			null, null, "'!='", "'=='", "'!'", "'||'", "'&&'", "'='", "'+='", "'-='", 
			"'>='", "'<='", "'>'", "'<'", "'*'", "'/'", "'+'", "'-'", "'%'", "'('", 
			"')'", "'{'", "'}'", "':'", "'?'", "'.'", "','", "';'", "'['", "']'", 
			"'_'", "'&'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "STRINGS", "INTS", "FLOATS", "BOOLS", "CHARACTERS", "VAR", "NIL", 
			"TRUE", "FALSE", "PRINT", "IF", "ELSE", "WHILE", "LET", "FOR", "IN", 
			"SWITCH", "CASE", "DEFAULT", "GUARD", "CONTINUE", "RETURN", "BREAK", 
			"APPEND", "REMOVELAST", "REMOVE", "AT", "ISEMPTY", "COUNT", "FUNC", "INOUT", 
			"STRUCT", "NUMBER", "STRING", "ID", "DIF", "IG_IG", "NOT", "OR", "AND", 
			"IG", "IG_ADD", "IG_SUB", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", 
			"DIV", "ADD", "SUB", "PERCENT", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", 
			"DOUBLEPTS", "QUESTION", "POINT", "COMA", "PTCOMA", "CORCHETEIZQ", "CORCHETEDER", 
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
			setState(84);
			((SContext)_localctx).block = block();
			setState(85);
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
			setState(89); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(88);
					((BlockContext)_localctx).instruction = instruction();
					((BlockContext)_localctx).ins.add(((BlockContext)_localctx).instruction);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(91); 
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
		public DeclarestructstmtContext declarestructstmt;
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
		public DeclarestructstmtContext declarestructstmt() {
			return getRuleContext(DeclarestructstmtContext.class,0);
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
			setState(152);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(95);
				((InstructionContext)_localctx).printstmt = printstmt();
				 _localctx.inst = ((InstructionContext)_localctx).printstmt.prnt
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(98);
				((InstructionContext)_localctx).declarestmt = declarestmt();
				_localctx.inst = ((InstructionContext)_localctx).declarestmt.dec
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(101);
				((InstructionContext)_localctx).constantstmt = constantstmt();
				_localctx.inst = ((InstructionContext)_localctx).constantstmt.const
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(104);
				((InstructionContext)_localctx).assignationstmt = assignationstmt();
				_localctx.inst = ((InstructionContext)_localctx).assignationstmt.assign
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(107);
				((InstructionContext)_localctx).increaseanddecreasestmt = increaseanddecreasestmt();
				_localctx.inst = ((InstructionContext)_localctx).increaseanddecreasestmt.increasedecrease
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(110);
				((InstructionContext)_localctx).ifstmt = ifstmt();
				 _localctx.inst = ((InstructionContext)_localctx).ifstmt.ift 
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(113);
				((InstructionContext)_localctx).whilestmt = whilestmt();
				_localctx.inst = ((InstructionContext)_localctx).whilestmt.while
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(116);
				((InstructionContext)_localctx).forstmt = forstmt();
				_localctx.inst = ((InstructionContext)_localctx).forstmt.for
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(119);
				((InstructionContext)_localctx).switchstmt = switchstmt();
				_localctx.inst = ((InstructionContext)_localctx).switchstmt.switch
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(122);
				((InstructionContext)_localctx).guardstmt = guardstmt();
				_localctx.inst = ((InstructionContext)_localctx).guardstmt.gua
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(125);
				((InstructionContext)_localctx).transferstmt = transferstmt();
				_localctx.inst = ((InstructionContext)_localctx).transferstmt.tran
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(128);
				((InstructionContext)_localctx).declarevectorstmt = declarevectorstmt();
				_localctx.inst = ((InstructionContext)_localctx).declarevectorstmt.decvec
				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(131);
				((InstructionContext)_localctx).appendstmt = appendstmt();
				_localctx.inst = ((InstructionContext)_localctx).appendstmt.app
				}
				break;
			case 14:
				enterOuterAlt(_localctx, 14);
				{
				setState(134);
				((InstructionContext)_localctx).removelaststmt = removelaststmt();
				_localctx.inst = ((InstructionContext)_localctx).removelaststmt.removl
				}
				break;
			case 15:
				enterOuterAlt(_localctx, 15);
				{
				setState(137);
				((InstructionContext)_localctx).removestmt = removestmt();
				_localctx.inst = ((InstructionContext)_localctx).removestmt.remov
				}
				break;
			case 16:
				enterOuterAlt(_localctx, 16);
				{
				setState(140);
				((InstructionContext)_localctx).assignationvecstmt = assignationvecstmt();
				_localctx.inst = ((InstructionContext)_localctx).assignationvecstmt.assignvec
				}
				break;
			case 17:
				enterOuterAlt(_localctx, 17);
				{
				setState(143);
				((InstructionContext)_localctx).declarefuncstmt = declarefuncstmt();
				_localctx.inst = ((InstructionContext)_localctx).declarefuncstmt.decfunc
				}
				break;
			case 18:
				enterOuterAlt(_localctx, 18);
				{
				setState(146);
				((InstructionContext)_localctx).accessfuncinstruction = accessfuncinstruction();
				_localctx.inst = ((InstructionContext)_localctx).accessfuncinstruction.accessfuncin
				}
				break;
			case 19:
				enterOuterAlt(_localctx, 19);
				{
				setState(149);
				((InstructionContext)_localctx).declarestructstmt = declarestructstmt();
				_localctx.inst = ((InstructionContext)_localctx).declarestructstmt.decstruct 
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
			setState(154);
			((PrintstmtContext)_localctx).PRINT = match(PRINT);
			setState(155);
			match(PARIZQ);
			setState(156);
			((PrintstmtContext)_localctx).listParams = listParams(0);
			setState(157);
			match(PARDER);
			setState(159);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PTCOMA) {
				{
				setState(158);
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
			setState(193);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(163);
				((DeclarestmtContext)_localctx).VAR = match(VAR);
				setState(164);
				((DeclarestmtContext)_localctx).ID = match(ID);
				setState(165);
				match(DOUBLEPTS);
				setState(166);
				((DeclarestmtContext)_localctx).type = type();
				setState(167);
				match(IG);
				setState(168);
				((DeclarestmtContext)_localctx).expr = expr(0);
				setState(170);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(169);
					match(PTCOMA);
					}
				}

				_localctx.dec = instructions.NewTodeclare((((DeclarestmtContext)_localctx).VAR!=null?((DeclarestmtContext)_localctx).VAR.getLine():0), (((DeclarestmtContext)_localctx).VAR!=null?((DeclarestmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarestmtContext)_localctx).ID!=null?((DeclarestmtContext)_localctx).ID.getText():null), ((DeclarestmtContext)_localctx).type.t, ((DeclarestmtContext)_localctx).expr.e, false)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(174);
				((DeclarestmtContext)_localctx).VAR = match(VAR);
				setState(175);
				((DeclarestmtContext)_localctx).ID = match(ID);
				setState(176);
				match(IG);
				setState(177);
				((DeclarestmtContext)_localctx).expr = expr(0);
				setState(179);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(178);
					match(PTCOMA);
					}
				}

				_localctx.dec = instructions.NewTodeclare((((DeclarestmtContext)_localctx).VAR!=null?((DeclarestmtContext)_localctx).VAR.getLine():0), (((DeclarestmtContext)_localctx).VAR!=null?((DeclarestmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarestmtContext)_localctx).ID!=null?((DeclarestmtContext)_localctx).ID.getText():null), environment.NULL, ((DeclarestmtContext)_localctx).expr.e, false)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(183);
				((DeclarestmtContext)_localctx).VAR = match(VAR);
				setState(184);
				((DeclarestmtContext)_localctx).ID = match(ID);
				setState(185);
				match(DOUBLEPTS);
				setState(186);
				((DeclarestmtContext)_localctx).type = type();
				setState(187);
				match(QUESTION);
				setState(189);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(188);
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
			setState(215);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(195);
				((ConstantstmtContext)_localctx).LET = match(LET);
				setState(196);
				((ConstantstmtContext)_localctx).ID = match(ID);
				setState(197);
				match(DOUBLEPTS);
				setState(198);
				((ConstantstmtContext)_localctx).type = type();
				setState(199);
				match(IG);
				setState(200);
				((ConstantstmtContext)_localctx).expr = expr(0);
				setState(202);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(201);
					match(PTCOMA);
					}
				}

				_localctx.const = instructions.NewTodeclare((((ConstantstmtContext)_localctx).LET!=null?((ConstantstmtContext)_localctx).LET.getLine():0), (((ConstantstmtContext)_localctx).LET!=null?((ConstantstmtContext)_localctx).LET.getCharPositionInLine():0), (((ConstantstmtContext)_localctx).ID!=null?((ConstantstmtContext)_localctx).ID.getText():null), ((ConstantstmtContext)_localctx).type.t, ((ConstantstmtContext)_localctx).expr.e, true)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(206);
				((ConstantstmtContext)_localctx).LET = match(LET);
				setState(207);
				((ConstantstmtContext)_localctx).ID = match(ID);
				setState(208);
				match(IG);
				setState(209);
				((ConstantstmtContext)_localctx).expr = expr(0);
				setState(211);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(210);
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
			setState(218); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(217);
					((BlockelifsContext)_localctx).ifstmt = ifstmt();
					((BlockelifsContext)_localctx).elif.add(((BlockelifsContext)_localctx).ifstmt);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(220); 
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
			setState(251);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(224);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(225);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(226);
				match(LLAVEIZQ);
				setState(227);
				((IfstmtContext)_localctx).ifb = block();
				setState(228);
				match(LLAVEDER);
				 ((IfstmtContext)_localctx).ift =  instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).ifb.blk, nil); 
				                                        _localctx.el = ((IfstmtContext)_localctx).ifb.blk
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(231);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(232);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(233);
				match(LLAVEIZQ);
				setState(234);
				((IfstmtContext)_localctx).ifelseblck = block();
				setState(235);
				match(LLAVEDER);
				setState(236);
				match(ELSE);
				setState(237);
				match(LLAVEIZQ);
				setState(238);
				((IfstmtContext)_localctx).elseifblck = block();
				setState(239);
				match(LLAVEDER);
				 ((IfstmtContext)_localctx).ift =  instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).ifelseblck.blk, ((IfstmtContext)_localctx).elseifblck.blk); 
				                                                                                        ((IfstmtContext)_localctx).el =  ((IfstmtContext)_localctx).ifelseblck.blk;
				                                                                                        
				                                                                                        
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(242);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(243);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(244);
				match(LLAVEIZQ);
				setState(245);
				((IfstmtContext)_localctx).elif = block();
				setState(246);
				match(LLAVEDER);
				setState(247);
				match(ELSE);
				setState(248);
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
			setState(269);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(253);
				((SwitchstmtContext)_localctx).SWITCH = match(SWITCH);
				setState(254);
				((SwitchstmtContext)_localctx).expr = expr(0);
				setState(255);
				match(LLAVEIZQ);
				setState(256);
				match(DEFAULT);
				setState(257);
				match(DOUBLEPTS);
				setState(258);
				((SwitchstmtContext)_localctx).block = block();
				setState(259);
				match(LLAVEDER);
				 ((SwitchstmtContext)_localctx).switch =  instructions.NewSwitch((((SwitchstmtContext)_localctx).SWITCH!=null?((SwitchstmtContext)_localctx).SWITCH.getLine():0), (((SwitchstmtContext)_localctx).SWITCH!=null?((SwitchstmtContext)_localctx).SWITCH.getCharPositionInLine():0), ((SwitchstmtContext)_localctx).expr.e, nil,  nil, ((SwitchstmtContext)_localctx).block.blk);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(262);
				((SwitchstmtContext)_localctx).SWITCH = match(SWITCH);
				setState(263);
				((SwitchstmtContext)_localctx).expr = expr(0);
				setState(264);
				match(LLAVEIZQ);
				setState(265);
				((SwitchstmtContext)_localctx).casestmt = casestmt();
				setState(266);
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
			setState(287);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(271);
				((CasestmtContext)_localctx).CASE = match(CASE);
				setState(272);
				((CasestmtContext)_localctx).expr = expr(0);
				setState(273);
				match(DOUBLEPTS);
				setState(274);
				((CasestmtContext)_localctx).block = block();
				setState(275);
				((CasestmtContext)_localctx).blockcases = blockcases();
				 _localctx.cas = instructions.NewCase((((CasestmtContext)_localctx).CASE!=null?((CasestmtContext)_localctx).CASE.getLine():0),(((CasestmtContext)_localctx).CASE!=null?((CasestmtContext)_localctx).CASE.getCharPositionInLine():0), ((CasestmtContext)_localctx).expr.e,((CasestmtContext)_localctx).block.blk, ((CasestmtContext)_localctx).blockcases.blkcase)
				                                fmt.Println("entro")
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(278);
				((CasestmtContext)_localctx).CASE = match(CASE);
				setState(279);
				((CasestmtContext)_localctx).expr = expr(0);
				setState(280);
				match(DOUBLEPTS);
				setState(281);
				((CasestmtContext)_localctx).sen = block();
				setState(282);
				match(DEFAULT);
				setState(283);
				match(DOUBLEPTS);
				setState(284);
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
			setState(290); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(289);
					((BlockcasesContext)_localctx).casestmt = casestmt();
					((BlockcasesContext)_localctx).casedef.add(((BlockcasesContext)_localctx).casestmt);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(292); 
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
			setState(296);
			((AssignationstmtContext)_localctx).ID = match(ID);
			setState(297);
			((AssignationstmtContext)_localctx).IG = match(IG);
			setState(298);
			((AssignationstmtContext)_localctx).expr = expr(0);
			setState(300);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PTCOMA) {
				{
				setState(299);
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
			setState(304);
			((WhilestmtContext)_localctx).WHILE = match(WHILE);
			setState(305);
			((WhilestmtContext)_localctx).expr = expr(0);
			setState(306);
			match(LLAVEIZQ);
			setState(307);
			((WhilestmtContext)_localctx).block = block();
			setState(308);
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
			setState(332);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(311);
				((ForstmtContext)_localctx).FOR = match(FOR);
				setState(312);
				((ForstmtContext)_localctx).ID = match(ID);
				setState(313);
				match(IN);
				setState(314);
				((ForstmtContext)_localctx).expr = expr(0);
				setState(315);
				match(LLAVEIZQ);
				setState(316);
				((ForstmtContext)_localctx).block = block();
				setState(317);
				match(LLAVEDER);
				_localctx.for = instructions.NewFor((((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getLine():0), (((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getCharPositionInLine():0), (((ForstmtContext)_localctx).ID!=null?((ForstmtContext)_localctx).ID.getText():null), ((ForstmtContext)_localctx).expr.e, nil, ((ForstmtContext)_localctx).block.blk)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(320);
				((ForstmtContext)_localctx).FOR = match(FOR);
				setState(321);
				((ForstmtContext)_localctx).ID = match(ID);
				setState(322);
				match(IN);
				setState(323);
				((ForstmtContext)_localctx).first = expr(0);
				setState(324);
				match(POINT);
				setState(325);
				match(POINT);
				setState(326);
				((ForstmtContext)_localctx).second = expr(0);
				setState(327);
				match(LLAVEIZQ);
				setState(328);
				((ForstmtContext)_localctx).block = block();
				setState(329);
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
			setState(334);
			((GuardstmtContext)_localctx).GUARD = match(GUARD);
			setState(335);
			((GuardstmtContext)_localctx).expr = expr(0);
			setState(336);
			match(ELSE);
			setState(337);
			match(LLAVEIZQ);
			setState(338);
			((GuardstmtContext)_localctx).block = block();
			setState(339);
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
			setState(340);
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
			setState(365);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,21,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(343);
				((TransferstmtContext)_localctx).RETURN = match(RETURN);
				setState(345);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(344);
					match(PTCOMA);
					}
				}

				_localctx.tran = instructions.NewReturnIn((((TransferstmtContext)_localctx).RETURN!=null?((TransferstmtContext)_localctx).RETURN.getLine():0), (((TransferstmtContext)_localctx).RETURN!=null?((TransferstmtContext)_localctx).RETURN.getCharPositionInLine():0), nil)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(348);
				((TransferstmtContext)_localctx).RETURN = match(RETURN);
				setState(349);
				((TransferstmtContext)_localctx).expr = expr(0);
				setState(351);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(350);
					match(PTCOMA);
					}
				}

				_localctx.tran = instructions.NewReturnIn((((TransferstmtContext)_localctx).RETURN!=null?((TransferstmtContext)_localctx).RETURN.getLine():0), (((TransferstmtContext)_localctx).RETURN!=null?((TransferstmtContext)_localctx).RETURN.getCharPositionInLine():0), ((TransferstmtContext)_localctx).expr.e)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(355);
				((TransferstmtContext)_localctx).CONTINUE = match(CONTINUE);
				setState(357);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(356);
					match(PTCOMA);
					}
				}

				_localctx.tran = instructions.NewContinue((((TransferstmtContext)_localctx).CONTINUE!=null?((TransferstmtContext)_localctx).CONTINUE.getLine():0), (((TransferstmtContext)_localctx).CONTINUE!=null?((TransferstmtContext)_localctx).CONTINUE.getCharPositionInLine():0))
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(360);
				((TransferstmtContext)_localctx).BREAK = match(BREAK);
				setState(362);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(361);
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
			setState(420);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,26,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(367);
				((DeclarevectorstmtContext)_localctx).VAR = match(VAR);
				setState(368);
				((DeclarevectorstmtContext)_localctx).ID = match(ID);
				setState(369);
				match(DOUBLEPTS);
				setState(370);
				match(CORCHETEIZQ);
				setState(371);
				((DeclarevectorstmtContext)_localctx).type = type();
				setState(372);
				match(CORCHETEDER);
				setState(373);
				match(IG);
				setState(374);
				match(CORCHETEIZQ);
				setState(375);
				((DeclarevectorstmtContext)_localctx).listParams = listParams(0);
				setState(376);
				match(CORCHETEDER);
				setState(378);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(377);
					match(PTCOMA);
					}
				}

				_localctx.decvec = instructions.NewToDeclareVector((((DeclarevectorstmtContext)_localctx).VAR!=null?((DeclarevectorstmtContext)_localctx).VAR.getLine():0), (((DeclarevectorstmtContext)_localctx).VAR!=null?((DeclarevectorstmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarevectorstmtContext)_localctx).ID!=null?((DeclarevectorstmtContext)_localctx).ID.getText():null), ((DeclarevectorstmtContext)_localctx).type.t, ((DeclarevectorstmtContext)_localctx).listParams.l, "")
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(382);
				((DeclarevectorstmtContext)_localctx).VAR = match(VAR);
				setState(383);
				((DeclarevectorstmtContext)_localctx).ID = match(ID);
				setState(384);
				match(DOUBLEPTS);
				setState(385);
				match(CORCHETEIZQ);
				setState(386);
				((DeclarevectorstmtContext)_localctx).type = type();
				setState(387);
				match(CORCHETEDER);
				setState(388);
				match(IG);
				setState(389);
				match(CORCHETEIZQ);
				setState(390);
				match(CORCHETEDER);
				setState(392);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(391);
					match(PTCOMA);
					}
				}

				_localctx.decvec = instructions.NewToDeclareVector((((DeclarevectorstmtContext)_localctx).VAR!=null?((DeclarevectorstmtContext)_localctx).VAR.getLine():0), (((DeclarevectorstmtContext)_localctx).VAR!=null?((DeclarevectorstmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarevectorstmtContext)_localctx).ID!=null?((DeclarevectorstmtContext)_localctx).ID.getText():null), ((DeclarevectorstmtContext)_localctx).type.t, nil, "")
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(396);
				match(VAR);
				setState(397);
				match(ID);
				setState(398);
				match(IG);
				setState(399);
				match(CORCHETEIZQ);
				setState(400);
				type();
				setState(401);
				match(CORCHETEDER);
				setState(402);
				match(PARIZQ);
				setState(403);
				match(PARDER);
				setState(405);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(404);
					match(PTCOMA);
					}
				}

				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(407);
				((DeclarevectorstmtContext)_localctx).VAR = match(VAR);
				setState(408);
				((DeclarevectorstmtContext)_localctx).F = match(ID);
				setState(409);
				match(DOUBLEPTS);
				setState(410);
				match(CORCHETEIZQ);
				setState(411);
				((DeclarevectorstmtContext)_localctx).type = type();
				setState(412);
				match(CORCHETEDER);
				setState(413);
				match(IG);
				setState(414);
				((DeclarevectorstmtContext)_localctx).S = match(ID);
				setState(416);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(415);
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
		public ListParamsContext listParams;
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode PARIZQ() { return getToken(GrammarParser.PARIZQ, 0); }
		public ListParamsContext listParams() {
			return getRuleContext(ListParamsContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(GrammarParser.PARDER, 0); }
		public TerminalNode PTCOMA() { return getToken(GrammarParser.PTCOMA, 0); }
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
			setState(422);
			((AccessfuncinstructionContext)_localctx).ID = match(ID);
			setState(423);
			match(PARIZQ);
			setState(424);
			((AccessfuncinstructionContext)_localctx).listParams = listParams(0);
			setState(425);
			match(PARDER);
			setState(426);
			match(PTCOMA);
			_localctx.accessfuncin = instructions.NewCallFunction((((AccessfuncinstructionContext)_localctx).ID!=null?((AccessfuncinstructionContext)_localctx).ID.getLine():0), (((AccessfuncinstructionContext)_localctx).ID!=null?((AccessfuncinstructionContext)_localctx).ID.getCharPositionInLine():0), (((AccessfuncinstructionContext)_localctx).ID!=null?((AccessfuncinstructionContext)_localctx).ID.getText():null), ((AccessfuncinstructionContext)_localctx).listParams.l, 1)
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
			setState(429);
			((AppendstmtContext)_localctx).ID = match(ID);
			setState(430);
			match(POINT);
			setState(431);
			match(APPEND);
			setState(432);
			match(PARIZQ);
			setState(433);
			((AppendstmtContext)_localctx).expr = expr(0);
			setState(434);
			match(PARDER);
			setState(436);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PTCOMA) {
				{
				setState(435);
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
			setState(440);
			((RemovelaststmtContext)_localctx).ID = match(ID);
			setState(441);
			match(POINT);
			setState(442);
			match(REMOVELAST);
			setState(443);
			match(PARIZQ);
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
			setState(450);
			((RemovestmtContext)_localctx).ID = match(ID);
			setState(451);
			match(POINT);
			setState(452);
			match(REMOVE);
			setState(453);
			match(PARIZQ);
			setState(454);
			match(AT);
			setState(455);
			match(DOUBLEPTS);
			setState(456);
			((RemovestmtContext)_localctx).expr = expr(0);
			setState(457);
			match(PARDER);
			setState(459);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PTCOMA) {
				{
				setState(458);
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
			setState(463);
			((EmptvecstmtContext)_localctx).ID = match(ID);
			setState(464);
			match(POINT);
			setState(465);
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
			setState(468);
			((CountvecstmtContext)_localctx).ID = match(ID);
			setState(469);
			match(POINT);
			setState(470);
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
			setState(473);
			((AccessvecstmtContext)_localctx).ID = match(ID);
			setState(474);
			match(CORCHETEIZQ);
			setState(475);
			((AccessvecstmtContext)_localctx).expr = expr(0);
			setState(476);
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
			setState(479);
			((AssignationvecstmtContext)_localctx).ID = match(ID);
			setState(480);
			match(CORCHETEIZQ);
			setState(481);
			((AssignationvecstmtContext)_localctx).expprim = expr(0);
			setState(482);
			match(CORCHETEDER);
			setState(483);
			match(IG);
			setState(484);
			((AssignationvecstmtContext)_localctx).expsegundo = expr(0);
			setState(486);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PTCOMA) {
				{
				setState(485);
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
			setState(490);
			match(VAR);
			setState(491);
			match(ID);
			setState(492);
			match(PARIZQ);
			setState(493);
			match(DOUBLEPTS);
			setState(494);
			type();
			setState(495);
			match(PARDER);
			setState(496);
			match(IG);
			setState(498);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PTCOMA) {
				{
				setState(497);
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
		public ListParamsFuncContext listParamsFunc;
		public TypeContext type;
		public BlockContext block;
		public TerminalNode FUNC() { return getToken(GrammarParser.FUNC, 0); }
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode PARIZQ() { return getToken(GrammarParser.PARIZQ, 0); }
		public ListParamsFuncContext listParamsFunc() {
			return getRuleContext(ListParamsFuncContext.class,0);
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
			setState(544);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,32,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(500);
				match(FUNC);
				setState(501);
				((DeclarefuncstmtContext)_localctx).ID = match(ID);
				setState(502);
				match(PARIZQ);
				setState(503);
				((DeclarefuncstmtContext)_localctx).listParamsFunc = listParamsFunc(0);
				setState(504);
				match(PARDER);
				setState(505);
				match(SUB);
				setState(506);
				match(MAYOR);
				setState(507);
				((DeclarefuncstmtContext)_localctx).type = type();
				setState(508);
				match(LLAVEIZQ);
				setState(509);
				((DeclarefuncstmtContext)_localctx).block = block();
				setState(510);
				match(LLAVEDER);
				_localctx.decfunc = instructions.NewToDeclareFunction((((DeclarefuncstmtContext)_localctx).ID!=null?((DeclarefuncstmtContext)_localctx).ID.getLine():0), (((DeclarefuncstmtContext)_localctx).ID!=null?((DeclarefuncstmtContext)_localctx).ID.getCharPositionInLine():0), (((DeclarefuncstmtContext)_localctx).ID!=null?((DeclarefuncstmtContext)_localctx).ID.getText():null), ((DeclarefuncstmtContext)_localctx).listParamsFunc.lf, ((DeclarefuncstmtContext)_localctx).type.t, ((DeclarefuncstmtContext)_localctx).block.blk, 1)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(513);
				match(FUNC);
				setState(514);
				((DeclarefuncstmtContext)_localctx).ID = match(ID);
				setState(515);
				match(PARIZQ);
				setState(516);
				((DeclarefuncstmtContext)_localctx).listParamsFunc = listParamsFunc(0);
				setState(517);
				match(PARDER);
				setState(518);
				match(LLAVEIZQ);
				setState(519);
				((DeclarefuncstmtContext)_localctx).block = block();
				setState(520);
				match(LLAVEDER);

				    
				    ((DeclarefuncstmtContext)_localctx).decfunc =  instructions.NewToDeclareFunction((((DeclarefuncstmtContext)_localctx).ID!=null?((DeclarefuncstmtContext)_localctx).ID.getLine():0), (((DeclarefuncstmtContext)_localctx).ID!=null?((DeclarefuncstmtContext)_localctx).ID.getCharPositionInLine():0), (((DeclarefuncstmtContext)_localctx).ID!=null?((DeclarefuncstmtContext)_localctx).ID.getText():null), ((DeclarefuncstmtContext)_localctx).listParamsFunc.lf, environment.NULL, ((DeclarefuncstmtContext)_localctx).block.blk, 2);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(523);
				match(FUNC);
				setState(524);
				((DeclarefuncstmtContext)_localctx).ID = match(ID);
				setState(525);
				match(PARIZQ);
				setState(526);
				match(PARDER);
				setState(527);
				match(SUB);
				setState(528);
				match(MAYOR);
				setState(529);
				((DeclarefuncstmtContext)_localctx).type = type();
				setState(530);
				match(LLAVEIZQ);
				setState(531);
				((DeclarefuncstmtContext)_localctx).block = block();
				setState(532);
				match(LLAVEDER);
				_localctx.decfunc = instructions.NewToDeclareFunction((((DeclarefuncstmtContext)_localctx).ID!=null?((DeclarefuncstmtContext)_localctx).ID.getLine():0), (((DeclarefuncstmtContext)_localctx).ID!=null?((DeclarefuncstmtContext)_localctx).ID.getCharPositionInLine():0), (((DeclarefuncstmtContext)_localctx).ID!=null?((DeclarefuncstmtContext)_localctx).ID.getText():null), nil, ((DeclarefuncstmtContext)_localctx).type.t, ((DeclarefuncstmtContext)_localctx).block.blk, 3)
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(535);
				match(FUNC);
				setState(536);
				((DeclarefuncstmtContext)_localctx).ID = match(ID);
				setState(537);
				match(PARIZQ);
				setState(538);
				match(PARDER);
				setState(539);
				match(LLAVEIZQ);
				setState(540);
				((DeclarefuncstmtContext)_localctx).block = block();
				setState(541);
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
			setState(547);
			((ListParamsFuncContext)_localctx).parameterfuncstmt = parameterfuncstmt();

			    
			            _localctx.lf = []interface{}{}
			            _localctx.lf = append(_localctx.lf, ((ListParamsFuncContext)_localctx).parameterfuncstmt.parameterfunc)
			        
			}
			_ctx.stop = _input.LT(-1);
			setState(557);
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
					setState(550);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(551);
					match(COMA);
					setState(552);
					((ListParamsFuncContext)_localctx).parameterfuncstmt = parameterfuncstmt();

					                              
					                                          var arrf []interface{}
					                                          arrf = append(((ListParamsFuncContext)_localctx).listf.lf, ((ListParamsFuncContext)_localctx).parameterfuncstmt.parameterfunc)
					                                          _localctx.lf = arrf
					                                      
					}
					} 
				}
				setState(559);
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
		public Token ID;
		public TypeContext type;
		public List<TerminalNode> ID() { return getTokens(GrammarParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(GrammarParser.ID, i);
		}
		public TerminalNode DOUBLEPTS() { return getToken(GrammarParser.DOUBLEPTS, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode GUION_BAJO() { return getToken(GrammarParser.GUION_BAJO, 0); }
		public TerminalNode CORCHETEIZQ() { return getToken(GrammarParser.CORCHETEIZQ, 0); }
		public TerminalNode CORCHETEDER() { return getToken(GrammarParser.CORCHETEDER, 0); }
		public TerminalNode INOUT() { return getToken(GrammarParser.INOUT, 0); }
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
			setState(577);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,35,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(560);
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
				setState(561);
				((ParameterfuncstmtContext)_localctx).ID = match(ID);
				setState(562);
				match(DOUBLEPTS);
				setState(563);
				((ParameterfuncstmtContext)_localctx).type = type();
				_localctx.parameterfunc = expressions.NewParameters((((ParameterfuncstmtContext)_localctx).exte!=null?((ParameterfuncstmtContext)_localctx).exte.getLine():0), (((ParameterfuncstmtContext)_localctx).exte!=null?((ParameterfuncstmtContext)_localctx).exte.getCharPositionInLine():0), ((ParameterfuncstmtContext)_localctx).type.t, (((ParameterfuncstmtContext)_localctx).exte!=null?((ParameterfuncstmtContext)_localctx).exte.getText():null),  (((ParameterfuncstmtContext)_localctx).ID!=null?((ParameterfuncstmtContext)_localctx).ID.getText():null), 1)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(566);
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
				setState(567);
				((ParameterfuncstmtContext)_localctx).ID = match(ID);
				setState(568);
				match(DOUBLEPTS);
				setState(570);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==INOUT) {
					{
					setState(569);
					match(INOUT);
					}
				}

				setState(572);
				match(CORCHETEIZQ);
				setState(573);
				((ParameterfuncstmtContext)_localctx).type = type();
				setState(574);
				match(CORCHETEDER);

				   
				    ((ParameterfuncstmtContext)_localctx).parameterfunc =  expressions.NewParameters((((ParameterfuncstmtContext)_localctx).exte!=null?((ParameterfuncstmtContext)_localctx).exte.getLine():0), (((ParameterfuncstmtContext)_localctx).exte!=null?((ParameterfuncstmtContext)_localctx).exte.getCharPositionInLine():0), ((ParameterfuncstmtContext)_localctx).type.t, (((ParameterfuncstmtContext)_localctx).exte!=null?((ParameterfuncstmtContext)_localctx).exte.getText():null),  (((ParameterfuncstmtContext)_localctx).ID!=null?((ParameterfuncstmtContext)_localctx).ID.getText():null),2);
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

	public static class DeclarestructstmtContext extends ParserRuleContext {
		public interfaces.Instruction decstruct;
		public Token STRUCT;
		public Token ID;
		public ListStructContext listStruct;
		public TerminalNode STRUCT() { return getToken(GrammarParser.STRUCT, 0); }
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(GrammarParser.LLAVEIZQ, 0); }
		public ListStructContext listStruct() {
			return getRuleContext(ListStructContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(GrammarParser.LLAVEDER, 0); }
		public DeclarestructstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declarestructstmt; }
	}

	public final DeclarestructstmtContext declarestructstmt() throws RecognitionException {
		DeclarestructstmtContext _localctx = new DeclarestructstmtContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_declarestructstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(579);
			((DeclarestructstmtContext)_localctx).STRUCT = match(STRUCT);
			setState(580);
			((DeclarestructstmtContext)_localctx).ID = match(ID);
			setState(581);
			match(LLAVEIZQ);
			setState(582);
			((DeclarestructstmtContext)_localctx).listStruct = listStruct(0);
			setState(583);
			match(LLAVEDER);
			 _localctx.decstruct = instructions.NewToDeclareStruct((((DeclarestructstmtContext)_localctx).STRUCT!=null?((DeclarestructstmtContext)_localctx).STRUCT.getLine():0), (((DeclarestructstmtContext)_localctx).STRUCT!=null?((DeclarestructstmtContext)_localctx).STRUCT.getCharPositionInLine():0), (((DeclarestructstmtContext)_localctx).ID!=null?((DeclarestructstmtContext)_localctx).ID.getText():null), ((DeclarestructstmtContext)_localctx).listStruct.l) 
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

	public static class ListStructContext extends ParserRuleContext {
		public []interface{} l;
		public ListStructContext list;
		public Token ID;
		public TypeContext type;
		public TerminalNode VAR() { return getToken(GrammarParser.VAR, 0); }
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode DOUBLEPTS() { return getToken(GrammarParser.DOUBLEPTS, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode COMA() { return getToken(GrammarParser.COMA, 0); }
		public ListStructContext listStruct() {
			return getRuleContext(ListStructContext.class,0);
		}
		public ListStructContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listStruct; }
	}

	public final ListStructContext listStruct() throws RecognitionException {
		return listStruct(0);
	}

	private ListStructContext listStruct(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ListStructContext _localctx = new ListStructContext(_ctx, _parentState);
		ListStructContext _prevctx = _localctx;
		int _startState = 60;
		enterRecursionRule(_localctx, 60, RULE_listStruct, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(594);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,36,_ctx) ) {
			case 1:
				{
				setState(587);
				match(VAR);
				setState(588);
				((ListStructContext)_localctx).ID = match(ID);
				setState(589);
				match(DOUBLEPTS);
				setState(590);
				((ListStructContext)_localctx).type = type();

				                        var arr []interface{}
				                        newParams := environment.NewStructType((((ListStructContext)_localctx).ID!=null?((ListStructContext)_localctx).ID.getText():null), ((ListStructContext)_localctx).type.t)
				                        arr = append(arr, newParams)
				                        _localctx.l = arr

				}
				break;
			case 2:
				{
				 _localctx.l = []interface{}{} 
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(606);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,37,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ListStructContext(_parentctx, _parentState);
					_localctx.list = _prevctx;
					_localctx.list = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_listStruct);
					setState(596);
					if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
					setState(597);
					match(COMA);
					setState(598);
					match(VAR);
					setState(599);
					((ListStructContext)_localctx).ID = match(ID);
					setState(600);
					match(DOUBLEPTS);
					setState(601);
					((ListStructContext)_localctx).type = type();

					                                      var arr []interface{}
					                                      newParams := environment.NewStructType((((ListStructContext)_localctx).ID!=null?((ListStructContext)_localctx).ID.getText():null), ((ListStructContext)_localctx).type.t)
					                                      arr = append(((ListStructContext)_localctx).list.l, newParams)
					                                      _localctx.l = arr

					          
					}
					} 
				}
				setState(608);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,37,_ctx);
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
		public EmptvecstmtContext emptvecstmt;
		public CountvecstmtContext countvecstmt;
		public AccessvecstmtContext accessvecstmt;
		public IntfunctionstmtContext intfunctionstmt;
		public FloatfunctionstmtContext floatfunctionstmt;
		public StringfunctionstmtContext stringfunctionstmt;
		public AccessfuncstmtContext accessfuncstmt;
		public Token ID;
		public ListStructExpContext listStructExp;
		public AccessstructstmtContext accessstructstmt;
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
		public AccessfuncstmtContext accessfuncstmt() {
			return getRuleContext(AccessfuncstmtContext.class,0);
		}
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(GrammarParser.LLAVEIZQ, 0); }
		public ListStructExpContext listStructExp() {
			return getRuleContext(ListStructExpContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(GrammarParser.LLAVEDER, 0); }
		public AccessstructstmtContext accessstructstmt() {
			return getRuleContext(AccessstructstmtContext.class,0);
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
		int _startState = 62;
		enterRecursionRule(_localctx, 62, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(666);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,38,_ctx) ) {
			case 1:
				{
				setState(610);
				((ExprContext)_localctx).op = match(NOT);
				setState(611);
				((ExprContext)_localctx).left = ((ExprContext)_localctx).expr = expr(18);
				 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).left.e) 
				}
				break;
			case 2:
				{
				setState(614);
				((ExprContext)_localctx).op = match(SUB);
				setState(615);
				((ExprContext)_localctx).left = ((ExprContext)_localctx).expr = expr(17);
				 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e,  (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), nil)
				}
				break;
			case 3:
				{
				setState(618);
				match(PARIZQ);
				setState(619);
				((ExprContext)_localctx).expr = expr(0);
				setState(620);
				match(PARDER);
				 _localctx.e = ((ExprContext)_localctx).expr.e 
				}
				break;
			case 4:
				{
				setState(623);
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
				setState(625);
				((ExprContext)_localctx).STRING = match(STRING);

				        str := (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getText():null)
				        _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getLine():0), (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getCharPositionInLine():0), str[1:len(str)-1],environment.STRING)
				    
				}
				break;
			case 6:
				{
				setState(627);
				((ExprContext)_localctx).TRUE = match(TRUE);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).TRUE!=null?((ExprContext)_localctx).TRUE.getLine():0), (((ExprContext)_localctx).TRUE!=null?((ExprContext)_localctx).TRUE.getCharPositionInLine():0), true, environment.BOOLEAN) 
				}
				break;
			case 7:
				{
				setState(629);
				((ExprContext)_localctx).FALSE = match(FALSE);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).FALSE!=null?((ExprContext)_localctx).FALSE.getLine():0), (((ExprContext)_localctx).FALSE!=null?((ExprContext)_localctx).FALSE.getCharPositionInLine():0), false, environment.BOOLEAN) 
				}
				break;
			case 8:
				{
				setState(631);
				((ExprContext)_localctx).NIL = match(NIL);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).NIL!=null?((ExprContext)_localctx).NIL.getLine():0), (((ExprContext)_localctx).NIL!=null?((ExprContext)_localctx).NIL.getCharPositionInLine():0), nil, environment.NULL) 
				}
				break;
			case 9:
				{
				setState(633);
				((ExprContext)_localctx).accessstmt = accessstmt();
				_localctx.e = ((ExprContext)_localctx).accessstmt.access
				}
				break;
			case 10:
				{
				setState(636);
				((ExprContext)_localctx).emptvecstmt = emptvecstmt();
				_localctx.e = ((ExprContext)_localctx).emptvecstmt.emptyvec
				}
				break;
			case 11:
				{
				setState(639);
				((ExprContext)_localctx).countvecstmt = countvecstmt();
				_localctx.e = ((ExprContext)_localctx).countvecstmt.count
				}
				break;
			case 12:
				{
				setState(642);
				((ExprContext)_localctx).accessvecstmt = accessvecstmt();
				_localctx.e = ((ExprContext)_localctx).accessvecstmt.accessvec
				}
				break;
			case 13:
				{
				setState(645);
				((ExprContext)_localctx).intfunctionstmt = intfunctionstmt();
				_localctx.e = ((ExprContext)_localctx).intfunctionstmt.intfunc
				}
				break;
			case 14:
				{
				setState(648);
				((ExprContext)_localctx).floatfunctionstmt = floatfunctionstmt();
				_localctx.e = ((ExprContext)_localctx).floatfunctionstmt.floatfunc
				}
				break;
			case 15:
				{
				setState(651);
				((ExprContext)_localctx).stringfunctionstmt = stringfunctionstmt();
				_localctx.e = ((ExprContext)_localctx).stringfunctionstmt.stringfunc
				}
				break;
			case 16:
				{
				setState(654);
				((ExprContext)_localctx).accessfuncstmt = accessfuncstmt();
				_localctx.e = ((ExprContext)_localctx).accessfuncstmt.funcexp
				}
				break;
			case 17:
				{
				setState(657);
				((ExprContext)_localctx).ID = match(ID);
				setState(658);
				match(LLAVEIZQ);
				setState(659);
				((ExprContext)_localctx).listStructExp = listStructExp(0);
				setState(660);
				match(LLAVEDER);
				 _localctx.e = expressions.NewStructExp((((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getCharPositionInLine():0), (((ExprContext)_localctx).ID!=null?((ExprContext)_localctx).ID.getText():null), ((ExprContext)_localctx).listStructExp.l ) 
				}
				break;
			case 18:
				{
				setState(663);
				((ExprContext)_localctx).accessstructstmt = accessstructstmt();
				_localctx.e = ((ExprContext)_localctx).accessstructstmt.accessstruct
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(705);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,40,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(703);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,39,_ctx) ) {
					case 1:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(668);
						if (!(precpred(_ctx, 25))) throw new FailedPredicateException(this, "precpred(_ctx, 25)");
						setState(669);
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
						setState(670);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(26);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 2:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(673);
						if (!(precpred(_ctx, 24))) throw new FailedPredicateException(this, "precpred(_ctx, 24)");
						setState(674);
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
						setState(675);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(25);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 3:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(678);
						if (!(precpred(_ctx, 23))) throw new FailedPredicateException(this, "precpred(_ctx, 23)");
						setState(679);
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
						setState(680);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(24);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 4:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(683);
						if (!(precpred(_ctx, 22))) throw new FailedPredicateException(this, "precpred(_ctx, 22)");
						setState(684);
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
						setState(685);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(23);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 5:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(688);
						if (!(precpred(_ctx, 21))) throw new FailedPredicateException(this, "precpred(_ctx, 21)");
						setState(689);
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
						setState(690);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(22);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 6:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(693);
						if (!(precpred(_ctx, 20))) throw new FailedPredicateException(this, "precpred(_ctx, 20)");
						setState(694);
						((ExprContext)_localctx).op = match(AND);
						setState(695);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(21);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 7:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(698);
						if (!(precpred(_ctx, 19))) throw new FailedPredicateException(this, "precpred(_ctx, 19)");
						setState(699);
						((ExprContext)_localctx).op = match(OR);
						setState(700);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(20);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					}
					} 
				}
				setState(707);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,40,_ctx);
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

	public static class AccessstructstmtContext extends ParserRuleContext {
		public interfaces.Expression accessstruct;
		public Token list;
		public Token ID;
		public TerminalNode POINT() { return getToken(GrammarParser.POINT, 0); }
		public List<TerminalNode> ID() { return getTokens(GrammarParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(GrammarParser.ID, i);
		}
		public AccessstructstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_accessstructstmt; }
	}

	public final AccessstructstmtContext accessstructstmt() throws RecognitionException {
		AccessstructstmtContext _localctx = new AccessstructstmtContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_accessstructstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(708);
			((AccessstructstmtContext)_localctx).list = match(ID);
			setState(709);
			match(POINT);
			setState(710);
			((AccessstructstmtContext)_localctx).ID = match(ID);
			 _localctx.accessstruct = expressions.NewAccessStruct((((AccessstructstmtContext)_localctx).list!=null?((AccessstructstmtContext)_localctx).list.getLine():0), (((AccessstructstmtContext)_localctx).list!=null?((AccessstructstmtContext)_localctx).list.getCharPositionInLine():0), (((AccessstructstmtContext)_localctx).list!=null?((AccessstructstmtContext)_localctx).list.getText():null), (((AccessstructstmtContext)_localctx).ID!=null?((AccessstructstmtContext)_localctx).ID.getText():null))  
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
		public interfaces.Expression funcexp;
		public Token ID;
		public ListParamsContext listParams;
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode PARIZQ() { return getToken(GrammarParser.PARIZQ, 0); }
		public ListParamsContext listParams() {
			return getRuleContext(ListParamsContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(GrammarParser.PARDER, 0); }
		public AccessfuncstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_accessfuncstmt; }
	}

	public final AccessfuncstmtContext accessfuncstmt() throws RecognitionException {
		AccessfuncstmtContext _localctx = new AccessfuncstmtContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_accessfuncstmt);
		try {
			setState(722);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,41,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(713);
				((AccessfuncstmtContext)_localctx).ID = match(ID);
				setState(714);
				match(PARIZQ);
				setState(715);
				((AccessfuncstmtContext)_localctx).listParams = listParams(0);
				setState(716);
				match(PARDER);
				_localctx.funcexp = expressions.NewAccessFunction((((AccessfuncstmtContext)_localctx).ID!=null?((AccessfuncstmtContext)_localctx).ID.getLine():0), (((AccessfuncstmtContext)_localctx).ID!=null?((AccessfuncstmtContext)_localctx).ID.getCharPositionInLine():0), (((AccessfuncstmtContext)_localctx).ID!=null?((AccessfuncstmtContext)_localctx).ID.getText():null), ((AccessfuncstmtContext)_localctx).listParams.l, 1)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(719);
				match(ID);
				setState(720);
				match(PARIZQ);
				setState(721);
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
		enterRule(_localctx, 68, RULE_intfunctionstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(724);
			((IntfunctionstmtContext)_localctx).INTS = match(INTS);
			setState(725);
			match(PARIZQ);
			setState(726);
			((IntfunctionstmtContext)_localctx).expr = expr(0);
			setState(727);
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
		enterRule(_localctx, 70, RULE_floatfunctionstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(730);
			((FloatfunctionstmtContext)_localctx).FLOATS = match(FLOATS);
			setState(731);
			match(PARIZQ);
			setState(732);
			((FloatfunctionstmtContext)_localctx).expr = expr(0);
			setState(733);
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
		enterRule(_localctx, 72, RULE_stringfunctionstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(736);
			((StringfunctionstmtContext)_localctx).STRINGS = match(STRINGS);
			setState(737);
			match(PARIZQ);
			setState(738);
			((StringfunctionstmtContext)_localctx).expr = expr(0);
			setState(739);
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
		enterRule(_localctx, 74, RULE_accessstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(742);
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
		enterRule(_localctx, 76, RULE_increaseanddecreasestmt);
		int _la;
		try {
			setState(761);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,44,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(745);
				((IncreaseanddecreasestmtContext)_localctx).ID = match(ID);
				setState(746);
				((IncreaseanddecreasestmtContext)_localctx).IG_ADD = match(IG_ADD);
				setState(747);
				((IncreaseanddecreasestmtContext)_localctx).expr = expr(0);
				setState(749);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(748);
					match(PTCOMA);
					}
				}

				_localctx.increasedecrease = instructions.NewIncreaseDecrease((((IncreaseanddecreasestmtContext)_localctx).ID!=null?((IncreaseanddecreasestmtContext)_localctx).ID.getLine():0), (((IncreaseanddecreasestmtContext)_localctx).ID!=null?((IncreaseanddecreasestmtContext)_localctx).ID.getCharPositionInLine():0), (((IncreaseanddecreasestmtContext)_localctx).ID!=null?((IncreaseanddecreasestmtContext)_localctx).ID.getText():null), (((IncreaseanddecreasestmtContext)_localctx).IG_ADD!=null?((IncreaseanddecreasestmtContext)_localctx).IG_ADD.getText():null), ((IncreaseanddecreasestmtContext)_localctx).expr.e)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(753);
				((IncreaseanddecreasestmtContext)_localctx).ID = match(ID);
				setState(754);
				((IncreaseanddecreasestmtContext)_localctx).IG_SUB = match(IG_SUB);
				setState(755);
				((IncreaseanddecreasestmtContext)_localctx).expr = expr(0);
				setState(757);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(756);
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
		enterRule(_localctx, 78, RULE_type);
		try {
			setState(773);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRINGS:
				enterOuterAlt(_localctx, 1);
				{
				setState(763);
				match(STRINGS);
				_localctx.t = environment.STRING
				}
				break;
			case INTS:
				enterOuterAlt(_localctx, 2);
				{
				setState(765);
				match(INTS);
				_localctx.t = environment.INTEGER
				}
				break;
			case FLOATS:
				enterOuterAlt(_localctx, 3);
				{
				setState(767);
				match(FLOATS);
				_localctx.t = environment.FLOAT
				}
				break;
			case BOOLS:
				enterOuterAlt(_localctx, 4);
				{
				setState(769);
				match(BOOLS);
				_localctx.t = environment.BOOLEAN 
				}
				break;
			case CHARACTERS:
				enterOuterAlt(_localctx, 5);
				{
				setState(771);
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
		public TerminalNode AND_SIMPLE() { return getToken(GrammarParser.AND_SIMPLE, 0); }
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
		int _startState = 80;
		enterRecursionRule(_localctx, 80, RULE_listParams, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(777);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AND_SIMPLE) {
				{
				setState(776);
				match(AND_SIMPLE);
				}
			}

			setState(779);
			((ListParamsContext)_localctx).expr = expr(0);

			            _localctx.l = []interface{}{}
			            _localctx.l = append(_localctx.l, ((ListParamsContext)_localctx).expr.e)
			        
			}
			_ctx.stop = _input.LT(-1);
			setState(792);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,48,_ctx);
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
					setState(782);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(783);
					match(COMA);
					setState(785);
					_errHandler.sync(this);
					_la = _input.LA(1);
					if (_la==AND_SIMPLE) {
						{
						setState(784);
						match(AND_SIMPLE);
						}
					}

					setState(787);
					((ListParamsContext)_localctx).expr = expr(0);

					                                          var arr []interface{}
					                                          arr = append(((ListParamsContext)_localctx).list.l, ((ListParamsContext)_localctx).expr.e)
					                                          _localctx.l = arr
					                                      
					}
					} 
				}
				setState(794);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,48,_ctx);
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

	public static class ListStructExpContext extends ParserRuleContext {
		public []interface{} l;
		public ListStructExpContext list;
		public Token ID;
		public ExprContext expr;
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode DOUBLEPTS() { return getToken(GrammarParser.DOUBLEPTS, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode COMA() { return getToken(GrammarParser.COMA, 0); }
		public ListStructExpContext listStructExp() {
			return getRuleContext(ListStructExpContext.class,0);
		}
		public ListStructExpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listStructExp; }
	}

	public final ListStructExpContext listStructExp() throws RecognitionException {
		return listStructExp(0);
	}

	private ListStructExpContext listStructExp(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ListStructExpContext _localctx = new ListStructExpContext(_ctx, _parentState);
		ListStructExpContext _prevctx = _localctx;
		int _startState = 82;
		enterRecursionRule(_localctx, 82, RULE_listStructExp, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(802);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,49,_ctx) ) {
			case 1:
				{
				setState(796);
				((ListStructExpContext)_localctx).ID = match(ID);
				setState(797);
				match(DOUBLEPTS);
				setState(798);
				((ListStructExpContext)_localctx).expr = expr(0);

				                    var arr []interface{}
				                    StrExp := environment.NewStructContent((((ListStructExpContext)_localctx).ID!=null?((ListStructExpContext)_localctx).ID.getText():null), ((ListStructExpContext)_localctx).expr.e)
				                    arr = append(arr, StrExp)
				                    _localctx.l = arr
				                
				}
				break;
			case 2:
				{

				        _localctx.l = []interface{}{}
				    
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(813);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,50,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ListStructExpContext(_parentctx, _parentState);
					_localctx.list = _prevctx;
					_localctx.list = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_listStructExp);
					setState(804);
					if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
					setState(805);
					match(COMA);
					setState(806);
					((ListStructExpContext)_localctx).ID = match(ID);
					setState(807);
					match(DOUBLEPTS);
					setState(808);
					((ListStructExpContext)_localctx).expr = expr(0);

					                                                      var arr []interface{}
					                                                      StrExp := environment.NewStructContent((((ListStructExpContext)_localctx).ID!=null?((ListStructExpContext)_localctx).ID.getText():null), ((ListStructExpContext)_localctx).expr.e)
					                                                      arr = append(((ListStructExpContext)_localctx).list.l, StrExp)
					                                                      _localctx.l = arr
					                                                  
					}
					} 
				}
				setState(815);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,50,_ctx);
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
		case 30:
			return listStruct_sempred((ListStructContext)_localctx, predIndex);
		case 31:
			return expr_sempred((ExprContext)_localctx, predIndex);
		case 40:
			return listParams_sempred((ListParamsContext)_localctx, predIndex);
		case 41:
			return listStructExp_sempred((ListStructExpContext)_localctx, predIndex);
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
	private boolean listStruct_sempred(ListStructContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return precpred(_ctx, 3);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 2:
			return precpred(_ctx, 25);
		case 3:
			return precpred(_ctx, 24);
		case 4:
			return precpred(_ctx, 23);
		case 5:
			return precpred(_ctx, 22);
		case 6:
			return precpred(_ctx, 21);
		case 7:
			return precpred(_ctx, 20);
		case 8:
			return precpred(_ctx, 19);
		}
		return true;
	}
	private boolean listParams_sempred(ListParamsContext _localctx, int predIndex) {
		switch (predIndex) {
		case 9:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean listStructExp_sempred(ListStructExpContext _localctx, int predIndex) {
		switch (predIndex) {
		case 10:
			return precpred(_ctx, 3);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3F\u0333\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\3"+
		"\2\3\2\3\2\3\2\3\3\6\3\\\n\3\r\3\16\3]\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4\u009b"+
		"\n\4\3\5\3\5\3\5\3\5\3\5\5\5\u00a2\n\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6"+
		"\3\6\5\6\u00ad\n\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6\u00b6\n\6\3\6\3\6\3"+
		"\6\3\6\3\6\3\6\3\6\3\6\5\6\u00c0\n\6\3\6\3\6\5\6\u00c4\n\6\3\7\3\7\3\7"+
		"\3\7\3\7\3\7\3\7\5\7\u00cd\n\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\5\7\u00d6\n"+
		"\7\3\7\3\7\5\7\u00da\n\7\3\b\6\b\u00dd\n\b\r\b\16\b\u00de\3\b\3\b\3\t"+
		"\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\5\t\u00fe\n\t\3\n\3\n\3\n\3\n\3\n\3"+
		"\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\5\n\u0110\n\n\3\13\3\13\3\13"+
		"\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\5\13"+
		"\u0122\n\13\3\f\6\f\u0125\n\f\r\f\16\f\u0126\3\f\3\f\3\r\3\r\3\r\3\r\5"+
		"\r\u012f\n\r\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\17\3\17\3\17"+
		"\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17"+
		"\3\17\3\17\3\17\3\17\5\17\u014f\n\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20"+
		"\3\20\3\20\3\21\3\21\5\21\u015c\n\21\3\21\3\21\3\21\3\21\5\21\u0162\n"+
		"\21\3\21\3\21\3\21\3\21\5\21\u0168\n\21\3\21\3\21\3\21\5\21\u016d\n\21"+
		"\3\21\5\21\u0170\n\21\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22"+
		"\3\22\5\22\u017d\n\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22"+
		"\3\22\3\22\5\22\u018b\n\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22"+
		"\3\22\3\22\5\22\u0198\n\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22"+
		"\5\22\u01a3\n\22\3\22\3\22\5\22\u01a7\n\22\3\23\3\23\3\23\3\23\3\23\3"+
		"\23\3\23\3\24\3\24\3\24\3\24\3\24\3\24\3\24\5\24\u01b7\n\24\3\24\3\24"+
		"\3\25\3\25\3\25\3\25\3\25\3\25\5\25\u01c1\n\25\3\25\3\25\3\26\3\26\3\26"+
		"\3\26\3\26\3\26\3\26\3\26\3\26\5\26\u01ce\n\26\3\26\3\26\3\27\3\27\3\27"+
		"\3\27\3\27\3\30\3\30\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\31\3\32"+
		"\3\32\3\32\3\32\3\32\3\32\3\32\5\32\u01e9\n\32\3\32\3\32\3\33\3\33\3\33"+
		"\3\33\3\33\3\33\3\33\3\33\5\33\u01f5\n\33\3\34\3\34\3\34\3\34\3\34\3\34"+
		"\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34"+
		"\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34"+
		"\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\5\34\u0223\n\34\3\35"+
		"\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\7\35\u022e\n\35\f\35\16\35\u0231"+
		"\13\35\3\36\3\36\3\36\3\36\3\36\3\36\3\36\3\36\3\36\3\36\5\36\u023d\n"+
		"\36\3\36\3\36\3\36\3\36\3\36\5\36\u0244\n\36\3\37\3\37\3\37\3\37\3\37"+
		"\3\37\3\37\3 \3 \3 \3 \3 \3 \3 \3 \5 \u0255\n \3 \3 \3 \3 \3 \3 \3 \3"+
		" \7 \u025f\n \f \16 \u0262\13 \3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!"+
		"\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!"+
		"\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\5!\u029d"+
		"\n!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!"+
		"\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\7!\u02c2\n!\f!\16!\u02c5\13!\3"+
		"\"\3\"\3\"\3\"\3\"\3#\3#\3#\3#\3#\3#\3#\3#\3#\5#\u02d5\n#\3$\3$\3$\3$"+
		"\3$\3$\3%\3%\3%\3%\3%\3%\3&\3&\3&\3&\3&\3&\3\'\3\'\3\'\3(\3(\3(\3(\5("+
		"\u02f0\n(\3(\3(\3(\3(\3(\3(\5(\u02f8\n(\3(\3(\5(\u02fc\n(\3)\3)\3)\3)"+
		"\3)\3)\3)\3)\3)\3)\5)\u0308\n)\3*\3*\5*\u030c\n*\3*\3*\3*\3*\3*\3*\5*"+
		"\u0314\n*\3*\3*\3*\7*\u0319\n*\f*\16*\u031c\13*\3+\3+\3+\3+\3+\3+\3+\5"+
		"+\u0325\n+\3+\3+\3+\3+\3+\3+\3+\7+\u032e\n+\f+\16+\u0331\13+\3+\2\78>"+
		"@RT,\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$&(*,.\60\62\64\668:<>@"+
		"BDFHJLNPRT\2\t\3\2\27\31\4\2%%BB\3\2\62\63\3\2\64\66\4\2..\60\60\4\2/"+
		"/\61\61\3\2&\'\2\u036c\2V\3\2\2\2\4[\3\2\2\2\6\u009a\3\2\2\2\b\u009c\3"+
		"\2\2\2\n\u00c3\3\2\2\2\f\u00d9\3\2\2\2\16\u00dc\3\2\2\2\20\u00fd\3\2\2"+
		"\2\22\u010f\3\2\2\2\24\u0121\3\2\2\2\26\u0124\3\2\2\2\30\u012a\3\2\2\2"+
		"\32\u0132\3\2\2\2\34\u014e\3\2\2\2\36\u0150\3\2\2\2 \u016f\3\2\2\2\"\u01a6"+
		"\3\2\2\2$\u01a8\3\2\2\2&\u01af\3\2\2\2(\u01ba\3\2\2\2*\u01c4\3\2\2\2,"+
		"\u01d1\3\2\2\2.\u01d6\3\2\2\2\60\u01db\3\2\2\2\62\u01e1\3\2\2\2\64\u01ec"+
		"\3\2\2\2\66\u0222\3\2\2\28\u0224\3\2\2\2:\u0243\3\2\2\2<\u0245\3\2\2\2"+
		">\u0254\3\2\2\2@\u029c\3\2\2\2B\u02c6\3\2\2\2D\u02d4\3\2\2\2F\u02d6\3"+
		"\2\2\2H\u02dc\3\2\2\2J\u02e2\3\2\2\2L\u02e8\3\2\2\2N\u02fb\3\2\2\2P\u0307"+
		"\3\2\2\2R\u0309\3\2\2\2T\u0324\3\2\2\2VW\5\4\3\2WX\7\2\2\3XY\b\2\1\2Y"+
		"\3\3\2\2\2Z\\\5\6\4\2[Z\3\2\2\2\\]\3\2\2\2][\3\2\2\2]^\3\2\2\2^_\3\2\2"+
		"\2_`\b\3\1\2`\5\3\2\2\2ab\5\b\5\2bc\b\4\1\2c\u009b\3\2\2\2de\5\n\6\2e"+
		"f\b\4\1\2f\u009b\3\2\2\2gh\5\f\7\2hi\b\4\1\2i\u009b\3\2\2\2jk\5\30\r\2"+
		"kl\b\4\1\2l\u009b\3\2\2\2mn\5N(\2no\b\4\1\2o\u009b\3\2\2\2pq\5\20\t\2"+
		"qr\b\4\1\2r\u009b\3\2\2\2st\5\32\16\2tu\b\4\1\2u\u009b\3\2\2\2vw\5\34"+
		"\17\2wx\b\4\1\2x\u009b\3\2\2\2yz\5\22\n\2z{\b\4\1\2{\u009b\3\2\2\2|}\5"+
		"\36\20\2}~\b\4\1\2~\u009b\3\2\2\2\177\u0080\5 \21\2\u0080\u0081\b\4\1"+
		"\2\u0081\u009b\3\2\2\2\u0082\u0083\5\"\22\2\u0083\u0084\b\4\1\2\u0084"+
		"\u009b\3\2\2\2\u0085\u0086\5&\24\2\u0086\u0087\b\4\1\2\u0087\u009b\3\2"+
		"\2\2\u0088\u0089\5(\25\2\u0089\u008a\b\4\1\2\u008a\u009b\3\2\2\2\u008b"+
		"\u008c\5*\26\2\u008c\u008d\b\4\1\2\u008d\u009b\3\2\2\2\u008e\u008f\5\62"+
		"\32\2\u008f\u0090\b\4\1\2\u0090\u009b\3\2\2\2\u0091\u0092\5\66\34\2\u0092"+
		"\u0093\b\4\1\2\u0093\u009b\3\2\2\2\u0094\u0095\5$\23\2\u0095\u0096\b\4"+
		"\1\2\u0096\u009b\3\2\2\2\u0097\u0098\5<\37\2\u0098\u0099\b\4\1\2\u0099"+
		"\u009b\3\2\2\2\u009aa\3\2\2\2\u009ad\3\2\2\2\u009ag\3\2\2\2\u009aj\3\2"+
		"\2\2\u009am\3\2\2\2\u009ap\3\2\2\2\u009as\3\2\2\2\u009av\3\2\2\2\u009a"+
		"y\3\2\2\2\u009a|\3\2\2\2\u009a\177\3\2\2\2\u009a\u0082\3\2\2\2\u009a\u0085"+
		"\3\2\2\2\u009a\u0088\3\2\2\2\u009a\u008b\3\2\2\2\u009a\u008e\3\2\2\2\u009a"+
		"\u0091\3\2\2\2\u009a\u0094\3\2\2\2\u009a\u0097\3\2\2\2\u009b\7\3\2\2\2"+
		"\u009c\u009d\7\f\2\2\u009d\u009e\7\67\2\2\u009e\u009f\5R*\2\u009f\u00a1"+
		"\78\2\2\u00a0\u00a2\7?\2\2\u00a1\u00a0\3\2\2\2\u00a1\u00a2\3\2\2\2\u00a2"+
		"\u00a3\3\2\2\2\u00a3\u00a4\b\5\1\2\u00a4\t\3\2\2\2\u00a5\u00a6\7\b\2\2"+
		"\u00a6\u00a7\7%\2\2\u00a7\u00a8\7;\2\2\u00a8\u00a9\5P)\2\u00a9\u00aa\7"+
		"+\2\2\u00aa\u00ac\5@!\2\u00ab\u00ad\7?\2\2\u00ac\u00ab\3\2\2\2\u00ac\u00ad"+
		"\3\2\2\2\u00ad\u00ae\3\2\2\2\u00ae\u00af\b\6\1\2\u00af\u00c4\3\2\2\2\u00b0"+
		"\u00b1\7\b\2\2\u00b1\u00b2\7%\2\2\u00b2\u00b3\7+\2\2\u00b3\u00b5\5@!\2"+
		"\u00b4\u00b6\7?\2\2\u00b5\u00b4\3\2\2\2\u00b5\u00b6\3\2\2\2\u00b6\u00b7"+
		"\3\2\2\2\u00b7\u00b8\b\6\1\2\u00b8\u00c4\3\2\2\2\u00b9\u00ba\7\b\2\2\u00ba"+
		"\u00bb\7%\2\2\u00bb\u00bc\7;\2\2\u00bc\u00bd\5P)\2\u00bd\u00bf\7<\2\2"+
		"\u00be\u00c0\7?\2\2\u00bf\u00be\3\2\2\2\u00bf\u00c0\3\2\2\2\u00c0\u00c1"+
		"\3\2\2\2\u00c1\u00c2\b\6\1\2\u00c2\u00c4\3\2\2\2\u00c3\u00a5\3\2\2\2\u00c3"+
		"\u00b0\3\2\2\2\u00c3\u00b9\3\2\2\2\u00c4\13\3\2\2\2\u00c5\u00c6\7\20\2"+
		"\2\u00c6\u00c7\7%\2\2\u00c7\u00c8\7;\2\2\u00c8\u00c9\5P)\2\u00c9\u00ca"+
		"\7+\2\2\u00ca\u00cc\5@!\2\u00cb\u00cd\7?\2\2\u00cc\u00cb\3\2\2\2\u00cc"+
		"\u00cd\3\2\2\2\u00cd\u00ce\3\2\2\2\u00ce\u00cf\b\7\1\2\u00cf\u00da\3\2"+
		"\2\2\u00d0\u00d1\7\20\2\2\u00d1\u00d2\7%\2\2\u00d2\u00d3\7+\2\2\u00d3"+
		"\u00d5\5@!\2\u00d4\u00d6\7?\2\2\u00d5\u00d4\3\2\2\2\u00d5\u00d6\3\2\2"+
		"\2\u00d6\u00d7\3\2\2\2\u00d7\u00d8\b\7\1\2\u00d8\u00da\3\2\2\2\u00d9\u00c5"+
		"\3\2\2\2\u00d9\u00d0\3\2\2\2\u00da\r\3\2\2\2\u00db\u00dd\5\20\t\2\u00dc"+
		"\u00db\3\2\2\2\u00dd\u00de\3\2\2\2\u00de\u00dc\3\2\2\2\u00de\u00df\3\2"+
		"\2\2\u00df\u00e0\3\2\2\2\u00e0\u00e1\b\b\1\2\u00e1\17\3\2\2\2\u00e2\u00e3"+
		"\7\r\2\2\u00e3\u00e4\5@!\2\u00e4\u00e5\79\2\2\u00e5\u00e6\5\4\3\2\u00e6"+
		"\u00e7\7:\2\2\u00e7\u00e8\b\t\1\2\u00e8\u00fe\3\2\2\2\u00e9\u00ea\7\r"+
		"\2\2\u00ea\u00eb\5@!\2\u00eb\u00ec\79\2\2\u00ec\u00ed\5\4\3\2\u00ed\u00ee"+
		"\7:\2\2\u00ee\u00ef\7\16\2\2\u00ef\u00f0\79\2\2\u00f0\u00f1\5\4\3\2\u00f1"+
		"\u00f2\7:\2\2\u00f2\u00f3\b\t\1\2\u00f3\u00fe\3\2\2\2\u00f4\u00f5\7\r"+
		"\2\2\u00f5\u00f6\5@!\2\u00f6\u00f7\79\2\2\u00f7\u00f8\5\4\3\2\u00f8\u00f9"+
		"\7:\2\2\u00f9\u00fa\7\16\2\2\u00fa\u00fb\5\16\b\2\u00fb\u00fc\b\t\1\2"+
		"\u00fc\u00fe\3\2\2\2\u00fd\u00e2\3\2\2\2\u00fd\u00e9\3\2\2\2\u00fd\u00f4"+
		"\3\2\2\2\u00fe\21\3\2\2\2\u00ff\u0100\7\23\2\2\u0100\u0101\5@!\2\u0101"+
		"\u0102\79\2\2\u0102\u0103\7\25\2\2\u0103\u0104\7;\2\2\u0104\u0105\5\4"+
		"\3\2\u0105\u0106\7:\2\2\u0106\u0107\b\n\1\2\u0107\u0110\3\2\2\2\u0108"+
		"\u0109\7\23\2\2\u0109\u010a\5@!\2\u010a\u010b\79\2\2\u010b\u010c\5\24"+
		"\13\2\u010c\u010d\7:\2\2\u010d\u010e\b\n\1\2\u010e\u0110\3\2\2\2\u010f"+
		"\u00ff\3\2\2\2\u010f\u0108\3\2\2\2\u0110\23\3\2\2\2\u0111\u0112\7\24\2"+
		"\2\u0112\u0113\5@!\2\u0113\u0114\7;\2\2\u0114\u0115\5\4\3\2\u0115\u0116"+
		"\5\26\f\2\u0116\u0117\b\13\1\2\u0117\u0122\3\2\2\2\u0118\u0119\7\24\2"+
		"\2\u0119\u011a\5@!\2\u011a\u011b\7;\2\2\u011b\u011c\5\4\3\2\u011c\u011d"+
		"\7\25\2\2\u011d\u011e\7;\2\2\u011e\u011f\5\4\3\2\u011f\u0120\b\13\1\2"+
		"\u0120\u0122\3\2\2\2\u0121\u0111\3\2\2\2\u0121\u0118\3\2\2\2\u0122\25"+
		"\3\2\2\2\u0123\u0125\5\24\13\2\u0124\u0123\3\2\2\2\u0125\u0126\3\2\2\2"+
		"\u0126\u0124\3\2\2\2\u0126\u0127\3\2\2\2\u0127\u0128\3\2\2\2\u0128\u0129"+
		"\b\f\1\2\u0129\27\3\2\2\2\u012a\u012b\7%\2\2\u012b\u012c\7+\2\2\u012c"+
		"\u012e\5@!\2\u012d\u012f\7?\2\2\u012e\u012d\3\2\2\2\u012e\u012f\3\2\2"+
		"\2\u012f\u0130\3\2\2\2\u0130\u0131\b\r\1\2\u0131\31\3\2\2\2\u0132\u0133"+
		"\7\17\2\2\u0133\u0134\5@!\2\u0134\u0135\79\2\2\u0135\u0136\5\4\3\2\u0136"+
		"\u0137\7:\2\2\u0137\u0138\b\16\1\2\u0138\33\3\2\2\2\u0139\u013a\7\21\2"+
		"\2\u013a\u013b\7%\2\2\u013b\u013c\7\22\2\2\u013c\u013d\5@!\2\u013d\u013e"+
		"\79\2\2\u013e\u013f\5\4\3\2\u013f\u0140\7:\2\2\u0140\u0141\b\17\1\2\u0141"+
		"\u014f\3\2\2\2\u0142\u0143\7\21\2\2\u0143\u0144\7%\2\2\u0144\u0145\7\22"+
		"\2\2\u0145\u0146\5@!\2\u0146\u0147\7=\2\2\u0147\u0148\7=\2\2\u0148\u0149"+
		"\5@!\2\u0149\u014a\79\2\2\u014a\u014b\5\4\3\2\u014b\u014c\7:\2\2\u014c"+
		"\u014d\b\17\1\2\u014d\u014f\3\2\2\2\u014e\u0139\3\2\2\2\u014e\u0142\3"+
		"\2\2\2\u014f\35\3\2\2\2\u0150\u0151\7\26\2\2\u0151\u0152\5@!\2\u0152\u0153"+
		"\7\16\2\2\u0153\u0154\79\2\2\u0154\u0155\5\4\3\2\u0155\u0156\t\2\2\2\u0156"+
		"\u0157\7:\2\2\u0157\u0158\b\20\1\2\u0158\37\3\2\2\2\u0159\u015b\7\30\2"+
		"\2\u015a\u015c\7?\2\2\u015b\u015a\3\2\2\2\u015b\u015c\3\2\2\2\u015c\u015d"+
		"\3\2\2\2\u015d\u0170\b\21\1\2\u015e\u015f\7\30\2\2\u015f\u0161\5@!\2\u0160"+
		"\u0162\7?\2\2\u0161\u0160\3\2\2\2\u0161\u0162\3\2\2\2\u0162\u0163\3\2"+
		"\2\2\u0163\u0164\b\21\1\2\u0164\u0170\3\2\2\2\u0165\u0167\7\27\2\2\u0166"+
		"\u0168\7?\2\2\u0167\u0166\3\2\2\2\u0167\u0168\3\2\2\2\u0168\u0169\3\2"+
		"\2\2\u0169\u0170\b\21\1\2\u016a\u016c\7\31\2\2\u016b\u016d\7?\2\2\u016c"+
		"\u016b\3\2\2\2\u016c\u016d\3\2\2\2\u016d\u016e\3\2\2\2\u016e\u0170\b\21"+
		"\1\2\u016f\u0159\3\2\2\2\u016f\u015e\3\2\2\2\u016f\u0165\3\2\2\2\u016f"+
		"\u016a\3\2\2\2\u0170!\3\2\2\2\u0171\u0172\7\b\2\2\u0172\u0173\7%\2\2\u0173"+
		"\u0174\7;\2\2\u0174\u0175\7@\2\2\u0175\u0176\5P)\2\u0176\u0177\7A\2\2"+
		"\u0177\u0178\7+\2\2\u0178\u0179\7@\2\2\u0179\u017a\5R*\2\u017a\u017c\7"+
		"A\2\2\u017b\u017d\7?\2\2\u017c\u017b\3\2\2\2\u017c\u017d\3\2\2\2\u017d"+
		"\u017e\3\2\2\2\u017e\u017f\b\22\1\2\u017f\u01a7\3\2\2\2\u0180\u0181\7"+
		"\b\2\2\u0181\u0182\7%\2\2\u0182\u0183\7;\2\2\u0183\u0184\7@\2\2\u0184"+
		"\u0185\5P)\2\u0185\u0186\7A\2\2\u0186\u0187\7+\2\2\u0187\u0188\7@\2\2"+
		"\u0188\u018a\7A\2\2\u0189\u018b\7?\2\2\u018a\u0189\3\2\2\2\u018a\u018b"+
		"\3\2\2\2\u018b\u018c\3\2\2\2\u018c\u018d\b\22\1\2\u018d\u01a7\3\2\2\2"+
		"\u018e\u018f\7\b\2\2\u018f\u0190\7%\2\2\u0190\u0191\7+\2\2\u0191\u0192"+
		"\7@\2\2\u0192\u0193\5P)\2\u0193\u0194\7A\2\2\u0194\u0195\7\67\2\2\u0195"+
		"\u0197\78\2\2\u0196\u0198\7?\2\2\u0197\u0196\3\2\2\2\u0197\u0198\3\2\2"+
		"\2\u0198\u01a7\3\2\2\2\u0199\u019a\7\b\2\2\u019a\u019b\7%\2\2\u019b\u019c"+
		"\7;\2\2\u019c\u019d\7@\2\2\u019d\u019e\5P)\2\u019e\u019f\7A\2\2\u019f"+
		"\u01a0\7+\2\2\u01a0\u01a2\7%\2\2\u01a1\u01a3\7?\2\2\u01a2\u01a1\3\2\2"+
		"\2\u01a2\u01a3\3\2\2\2\u01a3\u01a4\3\2\2\2\u01a4\u01a5\b\22\1\2\u01a5"+
		"\u01a7\3\2\2\2\u01a6\u0171\3\2\2\2\u01a6\u0180\3\2\2\2\u01a6\u018e\3\2"+
		"\2\2\u01a6\u0199\3\2\2\2\u01a7#\3\2\2\2\u01a8\u01a9\7%\2\2\u01a9\u01aa"+
		"\7\67\2\2\u01aa\u01ab\5R*\2\u01ab\u01ac\78\2\2\u01ac\u01ad\7?\2\2\u01ad"+
		"\u01ae\b\23\1\2\u01ae%\3\2\2\2\u01af\u01b0\7%\2\2\u01b0\u01b1\7=\2\2\u01b1"+
		"\u01b2\7\32\2\2\u01b2\u01b3\7\67\2\2\u01b3\u01b4\5@!\2\u01b4\u01b6\78"+
		"\2\2\u01b5\u01b7\7?\2\2\u01b6\u01b5\3\2\2\2\u01b6\u01b7\3\2\2\2\u01b7"+
		"\u01b8\3\2\2\2\u01b8\u01b9\b\24\1\2\u01b9\'\3\2\2\2\u01ba\u01bb\7%\2\2"+
		"\u01bb\u01bc\7=\2\2\u01bc\u01bd\7\33\2\2\u01bd\u01be\7\67\2\2\u01be\u01c0"+
		"\78\2\2\u01bf\u01c1\7?\2\2\u01c0\u01bf\3\2\2\2\u01c0\u01c1\3\2\2\2\u01c1"+
		"\u01c2\3\2\2\2\u01c2\u01c3\b\25\1\2\u01c3)\3\2\2\2\u01c4\u01c5\7%\2\2"+
		"\u01c5\u01c6\7=\2\2\u01c6\u01c7\7\34\2\2\u01c7\u01c8\7\67\2\2\u01c8\u01c9"+
		"\7\35\2\2\u01c9\u01ca\7;\2\2\u01ca\u01cb\5@!\2\u01cb\u01cd\78\2\2\u01cc"+
		"\u01ce\7?\2\2\u01cd\u01cc\3\2\2\2\u01cd\u01ce\3\2\2\2\u01ce\u01cf\3\2"+
		"\2\2\u01cf\u01d0\b\26\1\2\u01d0+\3\2\2\2\u01d1\u01d2\7%\2\2\u01d2\u01d3"+
		"\7=\2\2\u01d3\u01d4\7\36\2\2\u01d4\u01d5\b\27\1\2\u01d5-\3\2\2\2\u01d6"+
		"\u01d7\7%\2\2\u01d7\u01d8\7=\2\2\u01d8\u01d9\7\37\2\2\u01d9\u01da\b\30"+
		"\1\2\u01da/\3\2\2\2\u01db\u01dc\7%\2\2\u01dc\u01dd\7@\2\2\u01dd\u01de"+
		"\5@!\2\u01de\u01df\7A\2\2\u01df\u01e0\b\31\1\2\u01e0\61\3\2\2\2\u01e1"+
		"\u01e2\7%\2\2\u01e2\u01e3\7@\2\2\u01e3\u01e4\5@!\2\u01e4\u01e5\7A\2\2"+
		"\u01e5\u01e6\7+\2\2\u01e6\u01e8\5@!\2\u01e7\u01e9\7?\2\2\u01e8\u01e7\3"+
		"\2\2\2\u01e8\u01e9\3\2\2\2\u01e9\u01ea\3\2\2\2\u01ea\u01eb\b\32\1\2\u01eb"+
		"\63\3\2\2\2\u01ec\u01ed\7\b\2\2\u01ed\u01ee\7%\2\2\u01ee\u01ef\7\67\2"+
		"\2\u01ef\u01f0\7;\2\2\u01f0\u01f1\5P)\2\u01f1\u01f2\78\2\2\u01f2\u01f4"+
		"\7+\2\2\u01f3\u01f5\7?\2\2\u01f4\u01f3\3\2\2\2\u01f4\u01f5\3\2\2\2\u01f5"+
		"\65\3\2\2\2\u01f6\u01f7\7 \2\2\u01f7\u01f8\7%\2\2\u01f8\u01f9\7\67\2\2"+
		"\u01f9\u01fa\58\35\2\u01fa\u01fb\78\2\2\u01fb\u01fc\7\65\2\2\u01fc\u01fd"+
		"\7\60\2\2\u01fd\u01fe\5P)\2\u01fe\u01ff\79\2\2\u01ff\u0200\5\4\3\2\u0200"+
		"\u0201\7:\2\2\u0201\u0202\b\34\1\2\u0202\u0223\3\2\2\2\u0203\u0204\7 "+
		"\2\2\u0204\u0205\7%\2\2\u0205\u0206\7\67\2\2\u0206\u0207\58\35\2\u0207"+
		"\u0208\78\2\2\u0208\u0209\79\2\2\u0209\u020a\5\4\3\2\u020a\u020b\7:\2"+
		"\2\u020b\u020c\b\34\1\2\u020c\u0223\3\2\2\2\u020d\u020e\7 \2\2\u020e\u020f"+
		"\7%\2\2\u020f\u0210\7\67\2\2\u0210\u0211\78\2\2\u0211\u0212\7\65\2\2\u0212"+
		"\u0213\7\60\2\2\u0213\u0214\5P)\2\u0214\u0215\79\2\2\u0215\u0216\5\4\3"+
		"\2\u0216\u0217\7:\2\2\u0217\u0218\b\34\1\2\u0218\u0223\3\2\2\2\u0219\u021a"+
		"\7 \2\2\u021a\u021b\7%\2\2\u021b\u021c\7\67\2\2\u021c\u021d\78\2\2\u021d"+
		"\u021e\79\2\2\u021e\u021f\5\4\3\2\u021f\u0220\7:\2\2\u0220\u0221\b\34"+
		"\1\2\u0221\u0223\3\2\2\2\u0222\u01f6\3\2\2\2\u0222\u0203\3\2\2\2\u0222"+
		"\u020d\3\2\2\2\u0222\u0219\3\2\2\2\u0223\67\3\2\2\2\u0224\u0225\b\35\1"+
		"\2\u0225\u0226\5:\36\2\u0226\u0227\b\35\1\2\u0227\u022f\3\2\2\2\u0228"+
		"\u0229\f\4\2\2\u0229\u022a\7>\2\2\u022a\u022b\5:\36\2\u022b\u022c\b\35"+
		"\1\2\u022c\u022e\3\2\2\2\u022d\u0228\3\2\2\2\u022e\u0231\3\2\2\2\u022f"+
		"\u022d\3\2\2\2\u022f\u0230\3\2\2\2\u02309\3\2\2\2\u0231\u022f\3\2\2\2"+
		"\u0232\u0233\t\3\2\2\u0233\u0234\7%\2\2\u0234\u0235\7;\2\2\u0235\u0236"+
		"\5P)\2\u0236\u0237\b\36\1\2\u0237\u0244\3\2\2\2\u0238\u0239\t\3\2\2\u0239"+
		"\u023a\7%\2\2\u023a\u023c\7;\2\2\u023b\u023d\7!\2\2\u023c\u023b\3\2\2"+
		"\2\u023c\u023d\3\2\2\2\u023d\u023e\3\2\2\2\u023e\u023f\7@\2\2\u023f\u0240"+
		"\5P)\2\u0240\u0241\7A\2\2\u0241\u0242\b\36\1\2\u0242\u0244\3\2\2\2\u0243"+
		"\u0232\3\2\2\2\u0243\u0238\3\2\2\2\u0244;\3\2\2\2\u0245\u0246\7\"\2\2"+
		"\u0246\u0247\7%\2\2\u0247\u0248\79\2\2\u0248\u0249\5> \2\u0249\u024a\7"+
		":\2\2\u024a\u024b\b\37\1\2\u024b=\3\2\2\2\u024c\u024d\b \1\2\u024d\u024e"+
		"\7\b\2\2\u024e\u024f\7%\2\2\u024f\u0250\7;\2\2\u0250\u0251\5P)\2\u0251"+
		"\u0252\b \1\2\u0252\u0255\3\2\2\2\u0253\u0255\b \1\2\u0254\u024c\3\2\2"+
		"\2\u0254\u0253\3\2\2\2\u0255\u0260\3\2\2\2\u0256\u0257\f\5\2\2\u0257\u0258"+
		"\7>\2\2\u0258\u0259\7\b\2\2\u0259\u025a\7%\2\2\u025a\u025b\7;\2\2\u025b"+
		"\u025c\5P)\2\u025c\u025d\b \1\2\u025d\u025f\3\2\2\2\u025e\u0256\3\2\2"+
		"\2\u025f\u0262\3\2\2\2\u0260\u025e\3\2\2\2\u0260\u0261\3\2\2\2\u0261?"+
		"\3\2\2\2\u0262\u0260\3\2\2\2\u0263\u0264\b!\1\2\u0264\u0265\7(\2\2\u0265"+
		"\u0266\5@!\24\u0266\u0267\b!\1\2\u0267\u029d\3\2\2\2\u0268\u0269\7\65"+
		"\2\2\u0269\u026a\5@!\23\u026a\u026b\b!\1\2\u026b\u029d\3\2\2\2\u026c\u026d"+
		"\7\67\2\2\u026d\u026e\5@!\2\u026e\u026f\78\2\2\u026f\u0270\b!\1\2\u0270"+
		"\u029d\3\2\2\2\u0271\u0272\7#\2\2\u0272\u029d\b!\1\2\u0273\u0274\7$\2"+
		"\2\u0274\u029d\b!\1\2\u0275\u0276\7\n\2\2\u0276\u029d\b!\1\2\u0277\u0278"+
		"\7\13\2\2\u0278\u029d\b!\1\2\u0279\u027a\7\t\2\2\u027a\u029d\b!\1\2\u027b"+
		"\u027c\5L\'\2\u027c\u027d\b!\1\2\u027d\u029d\3\2\2\2\u027e\u027f\5,\27"+
		"\2\u027f\u0280\b!\1\2\u0280\u029d\3\2\2\2\u0281\u0282\5.\30\2\u0282\u0283"+
		"\b!\1\2\u0283\u029d\3\2\2\2\u0284\u0285\5\60\31\2\u0285\u0286\b!\1\2\u0286"+
		"\u029d\3\2\2\2\u0287\u0288\5F$\2\u0288\u0289\b!\1\2\u0289\u029d\3\2\2"+
		"\2\u028a\u028b\5H%\2\u028b\u028c\b!\1\2\u028c\u029d\3\2\2\2\u028d\u028e"+
		"\5J&\2\u028e\u028f\b!\1\2\u028f\u029d\3\2\2\2\u0290\u0291\5D#\2\u0291"+
		"\u0292\b!\1\2\u0292\u029d\3\2\2\2\u0293\u0294\7%\2\2\u0294\u0295\79\2"+
		"\2\u0295\u0296\5T+\2\u0296\u0297\7:\2\2\u0297\u0298\b!\1\2\u0298\u029d"+
		"\3\2\2\2\u0299\u029a\5B\"\2\u029a\u029b\b!\1\2\u029b\u029d\3\2\2\2\u029c"+
		"\u0263\3\2\2\2\u029c\u0268\3\2\2\2\u029c\u026c\3\2\2\2\u029c\u0271\3\2"+
		"\2\2\u029c\u0273\3\2\2\2\u029c\u0275\3\2\2\2\u029c\u0277\3\2\2\2\u029c"+
		"\u0279\3\2\2\2\u029c\u027b\3\2\2\2\u029c\u027e\3\2\2\2\u029c\u0281\3\2"+
		"\2\2\u029c\u0284\3\2\2\2\u029c\u0287\3\2\2\2\u029c\u028a\3\2\2\2\u029c"+
		"\u028d\3\2\2\2\u029c\u0290\3\2\2\2\u029c\u0293\3\2\2\2\u029c\u0299\3\2"+
		"\2\2\u029d\u02c3\3\2\2\2\u029e\u029f\f\33\2\2\u029f\u02a0\t\4\2\2\u02a0"+
		"\u02a1\5@!\34\u02a1\u02a2\b!\1\2\u02a2\u02c2\3\2\2\2\u02a3\u02a4\f\32"+
		"\2\2\u02a4\u02a5\t\5\2\2\u02a5\u02a6\5@!\33\u02a6\u02a7\b!\1\2\u02a7\u02c2"+
		"\3\2\2\2\u02a8\u02a9\f\31\2\2\u02a9\u02aa\t\6\2\2\u02aa\u02ab\5@!\32\u02ab"+
		"\u02ac\b!\1\2\u02ac\u02c2\3\2\2\2\u02ad\u02ae\f\30\2\2\u02ae\u02af\t\7"+
		"\2\2\u02af\u02b0\5@!\31\u02b0\u02b1\b!\1\2\u02b1\u02c2\3\2\2\2\u02b2\u02b3"+
		"\f\27\2\2\u02b3\u02b4\t\b\2\2\u02b4\u02b5\5@!\30\u02b5\u02b6\b!\1\2\u02b6"+
		"\u02c2\3\2\2\2\u02b7\u02b8\f\26\2\2\u02b8\u02b9\7*\2\2\u02b9\u02ba\5@"+
		"!\27\u02ba\u02bb\b!\1\2\u02bb\u02c2\3\2\2\2\u02bc\u02bd\f\25\2\2\u02bd"+
		"\u02be\7)\2\2\u02be\u02bf\5@!\26\u02bf\u02c0\b!\1\2\u02c0\u02c2\3\2\2"+
		"\2\u02c1\u029e\3\2\2\2\u02c1\u02a3\3\2\2\2\u02c1\u02a8\3\2\2\2\u02c1\u02ad"+
		"\3\2\2\2\u02c1\u02b2\3\2\2\2\u02c1\u02b7\3\2\2\2\u02c1\u02bc\3\2\2\2\u02c2"+
		"\u02c5\3\2\2\2\u02c3\u02c1\3\2\2\2\u02c3\u02c4\3\2\2\2\u02c4A\3\2\2\2"+
		"\u02c5\u02c3\3\2\2\2\u02c6\u02c7\7%\2\2\u02c7\u02c8\7=\2\2\u02c8\u02c9"+
		"\7%\2\2\u02c9\u02ca\b\"\1\2\u02caC\3\2\2\2\u02cb\u02cc\7%\2\2\u02cc\u02cd"+
		"\7\67\2\2\u02cd\u02ce\5R*\2\u02ce\u02cf\78\2\2\u02cf\u02d0\b#\1\2\u02d0"+
		"\u02d5\3\2\2\2\u02d1\u02d2\7%\2\2\u02d2\u02d3\7\67\2\2\u02d3\u02d5\78"+
		"\2\2\u02d4\u02cb\3\2\2\2\u02d4\u02d1\3\2\2\2\u02d5E\3\2\2\2\u02d6\u02d7"+
		"\7\4\2\2\u02d7\u02d8\7\67\2\2\u02d8\u02d9\5@!\2\u02d9\u02da\78\2\2\u02da"+
		"\u02db\b$\1\2\u02dbG\3\2\2\2\u02dc\u02dd\7\5\2\2\u02dd\u02de\7\67\2\2"+
		"\u02de\u02df\5@!\2\u02df\u02e0\78\2\2\u02e0\u02e1\b%\1\2\u02e1I\3\2\2"+
		"\2\u02e2\u02e3\7\3\2\2\u02e3\u02e4\7\67\2\2\u02e4\u02e5\5@!\2\u02e5\u02e6"+
		"\78\2\2\u02e6\u02e7\b&\1\2\u02e7K\3\2\2\2\u02e8\u02e9\7%\2\2\u02e9\u02ea"+
		"\b\'\1\2\u02eaM\3\2\2\2\u02eb\u02ec\7%\2\2\u02ec\u02ed\7,\2\2\u02ed\u02ef"+
		"\5@!\2\u02ee\u02f0\7?\2\2\u02ef\u02ee\3\2\2\2\u02ef\u02f0\3\2\2\2\u02f0"+
		"\u02f1\3\2\2\2\u02f1\u02f2\b(\1\2\u02f2\u02fc\3\2\2\2\u02f3\u02f4\7%\2"+
		"\2\u02f4\u02f5\7-\2\2\u02f5\u02f7\5@!\2\u02f6\u02f8\7?\2\2\u02f7\u02f6"+
		"\3\2\2\2\u02f7\u02f8\3\2\2\2\u02f8\u02f9\3\2\2\2\u02f9\u02fa\b(\1\2\u02fa"+
		"\u02fc\3\2\2\2\u02fb\u02eb\3\2\2\2\u02fb\u02f3\3\2\2\2\u02fcO\3\2\2\2"+
		"\u02fd\u02fe\7\3\2\2\u02fe\u0308\b)\1\2\u02ff\u0300\7\4\2\2\u0300\u0308"+
		"\b)\1\2\u0301\u0302\7\5\2\2\u0302\u0308\b)\1\2\u0303\u0304\7\6\2\2\u0304"+
		"\u0308\b)\1\2\u0305\u0306\7\7\2\2\u0306\u0308\b)\1\2\u0307\u02fd\3\2\2"+
		"\2\u0307\u02ff\3\2\2\2\u0307\u0301\3\2\2\2\u0307\u0303\3\2\2\2\u0307\u0305"+
		"\3\2\2\2\u0308Q\3\2\2\2\u0309\u030b\b*\1\2\u030a\u030c\7C\2\2\u030b\u030a"+
		"\3\2\2\2\u030b\u030c\3\2\2\2\u030c\u030d\3\2\2\2\u030d\u030e\5@!\2\u030e"+
		"\u030f\b*\1\2\u030f\u031a\3\2\2\2\u0310\u0311\f\4\2\2\u0311\u0313\7>\2"+
		"\2\u0312\u0314\7C\2\2\u0313\u0312\3\2\2\2\u0313\u0314\3\2\2\2\u0314\u0315"+
		"\3\2\2\2\u0315\u0316\5@!\2\u0316\u0317\b*\1\2\u0317\u0319\3\2\2\2\u0318"+
		"\u0310\3\2\2\2\u0319\u031c\3\2\2\2\u031a\u0318\3\2\2\2\u031a\u031b\3\2"+
		"\2\2\u031bS\3\2\2\2\u031c\u031a\3\2\2\2\u031d\u031e\b+\1\2\u031e\u031f"+
		"\7%\2\2\u031f\u0320\7;\2\2\u0320\u0321\5@!\2\u0321\u0322\b+\1\2\u0322"+
		"\u0325\3\2\2\2\u0323\u0325\b+\1\2\u0324\u031d\3\2\2\2\u0324\u0323\3\2"+
		"\2\2\u0325\u032f\3\2\2\2\u0326\u0327\f\5\2\2\u0327\u0328\7>\2\2\u0328"+
		"\u0329\7%\2\2\u0329\u032a\7;\2\2\u032a\u032b\5@!\2\u032b\u032c\b+\1\2"+
		"\u032c\u032e\3\2\2\2\u032d\u0326\3\2\2\2\u032e\u0331\3\2\2\2\u032f\u032d"+
		"\3\2\2\2\u032f\u0330\3\2\2\2\u0330U\3\2\2\2\u0331\u032f\3\2\2\2\65]\u009a"+
		"\u00a1\u00ac\u00b5\u00bf\u00c3\u00cc\u00d5\u00d9\u00de\u00fd\u010f\u0121"+
		"\u0126\u012e\u014e\u015b\u0161\u0167\u016c\u016f\u017c\u018a\u0197\u01a2"+
		"\u01a6\u01b6\u01c0\u01cd\u01e8\u01f4\u0222\u022f\u023c\u0243\u0254\u0260"+
		"\u029c\u02c1\u02c3\u02d4\u02ef\u02f7\u02fb\u0307\u030b\u0313\u031a\u0324"+
		"\u032f";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
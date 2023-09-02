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
		CORCHETEIZQ=61, CORCHETEDER=62, WHITESPACE=63, COMMENT=64, LINE_COMMENT=65;
	public static final int
		RULE_s = 0, RULE_block = 1, RULE_instruction = 2, RULE_printstmt = 3, 
		RULE_declarestmt = 4, RULE_constantstmt = 5, RULE_blockelifs = 6, RULE_ifstmt = 7, 
		RULE_switchstmt = 8, RULE_casestmt = 9, RULE_blockcases = 10, RULE_assignationstmt = 11, 
		RULE_whilestmt = 12, RULE_forstmt = 13, RULE_guardstmt = 14, RULE_transferstmt = 15, 
		RULE_declarevectorstmt = 16, RULE_appendstmt = 17, RULE_removelaststmt = 18, 
		RULE_removestmt = 19, RULE_emptvecstmt = 20, RULE_countvecstmt = 21, RULE_accessvecstmt = 22, 
		RULE_assignationvecstmt = 23, RULE_declarematrixstmt = 24, RULE_declarefuncstmt = 25, 
		RULE_expr = 26, RULE_accessstmt = 27, RULE_increaseanddecreasestmt = 28, 
		RULE_type = 29, RULE_listParams = 30;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "block", "instruction", "printstmt", "declarestmt", "constantstmt", 
			"blockelifs", "ifstmt", "switchstmt", "casestmt", "blockcases", "assignationstmt", 
			"whilestmt", "forstmt", "guardstmt", "transferstmt", "declarevectorstmt", 
			"appendstmt", "removelaststmt", "removestmt", "emptvecstmt", "countvecstmt", 
			"accessvecstmt", "assignationvecstmt", "declarematrixstmt", "declarefuncstmt", 
			"expr", "accessstmt", "increaseanddecreasestmt", "type", "listParams"
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
			"'{'", "'}'", "':'", "'?'", "'.'", "','", "';'", "'['", "']'"
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
			"WHITESPACE", "COMMENT", "LINE_COMMENT"
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
			setState(62);
			((SContext)_localctx).block = block();
			setState(63);
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
			setState(67); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(66);
					((BlockContext)_localctx).instruction = instruction();
					((BlockContext)_localctx).ins.add(((BlockContext)_localctx).instruction);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(69); 
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
		public InstructionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruction; }
	}

	public final InstructionContext instruction() throws RecognitionException {
		InstructionContext _localctx = new InstructionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_instruction);
		try {
			setState(124);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(73);
				((InstructionContext)_localctx).printstmt = printstmt();
				 _localctx.inst = ((InstructionContext)_localctx).printstmt.prnt
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(76);
				((InstructionContext)_localctx).declarestmt = declarestmt();
				_localctx.inst = ((InstructionContext)_localctx).declarestmt.dec
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(79);
				((InstructionContext)_localctx).constantstmt = constantstmt();
				_localctx.inst = ((InstructionContext)_localctx).constantstmt.const
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(82);
				((InstructionContext)_localctx).assignationstmt = assignationstmt();
				_localctx.inst = ((InstructionContext)_localctx).assignationstmt.assign
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(85);
				((InstructionContext)_localctx).increaseanddecreasestmt = increaseanddecreasestmt();
				_localctx.inst = ((InstructionContext)_localctx).increaseanddecreasestmt.increasedecrease
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(88);
				((InstructionContext)_localctx).ifstmt = ifstmt();
				 _localctx.inst = ((InstructionContext)_localctx).ifstmt.ift 
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(91);
				((InstructionContext)_localctx).whilestmt = whilestmt();
				_localctx.inst = ((InstructionContext)_localctx).whilestmt.while
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(94);
				((InstructionContext)_localctx).forstmt = forstmt();
				_localctx.inst = ((InstructionContext)_localctx).forstmt.for
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(97);
				((InstructionContext)_localctx).switchstmt = switchstmt();
				_localctx.inst = ((InstructionContext)_localctx).switchstmt.switch
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(100);
				((InstructionContext)_localctx).guardstmt = guardstmt();
				_localctx.inst = ((InstructionContext)_localctx).guardstmt.gua
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(103);
				((InstructionContext)_localctx).transferstmt = transferstmt();
				_localctx.inst = ((InstructionContext)_localctx).transferstmt.tran
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(106);
				((InstructionContext)_localctx).declarevectorstmt = declarevectorstmt();
				_localctx.inst = ((InstructionContext)_localctx).declarevectorstmt.decvec
				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(109);
				((InstructionContext)_localctx).appendstmt = appendstmt();
				_localctx.inst = ((InstructionContext)_localctx).appendstmt.app
				}
				break;
			case 14:
				enterOuterAlt(_localctx, 14);
				{
				setState(112);
				((InstructionContext)_localctx).removelaststmt = removelaststmt();
				_localctx.inst = ((InstructionContext)_localctx).removelaststmt.removl
				}
				break;
			case 15:
				enterOuterAlt(_localctx, 15);
				{
				setState(115);
				((InstructionContext)_localctx).removestmt = removestmt();
				_localctx.inst = ((InstructionContext)_localctx).removestmt.remov
				}
				break;
			case 16:
				enterOuterAlt(_localctx, 16);
				{
				setState(118);
				((InstructionContext)_localctx).assignationvecstmt = assignationvecstmt();
				_localctx.inst = ((InstructionContext)_localctx).assignationvecstmt.assignvec
				}
				break;
			case 17:
				enterOuterAlt(_localctx, 17);
				{
				setState(121);
				((InstructionContext)_localctx).declarefuncstmt = declarefuncstmt();
				_localctx.inst = ((InstructionContext)_localctx).declarefuncstmt.decfunc
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
			setState(126);
			((PrintstmtContext)_localctx).PRINT = match(PRINT);
			setState(127);
			match(PARIZQ);
			setState(128);
			((PrintstmtContext)_localctx).expr = expr(0);
			setState(129);
			match(PARDER);
			setState(131);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PTCOMA) {
				{
				setState(130);
				match(PTCOMA);
				}
			}

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
			setState(165);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(135);
				((DeclarestmtContext)_localctx).VAR = match(VAR);
				setState(136);
				((DeclarestmtContext)_localctx).ID = match(ID);
				setState(137);
				match(DOUBLEPTS);
				setState(138);
				((DeclarestmtContext)_localctx).type = type();
				setState(139);
				match(IG);
				setState(140);
				((DeclarestmtContext)_localctx).expr = expr(0);
				setState(142);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(141);
					match(PTCOMA);
					}
				}

				_localctx.dec = instructions.NewTodeclare((((DeclarestmtContext)_localctx).VAR!=null?((DeclarestmtContext)_localctx).VAR.getLine():0), (((DeclarestmtContext)_localctx).VAR!=null?((DeclarestmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarestmtContext)_localctx).ID!=null?((DeclarestmtContext)_localctx).ID.getText():null), ((DeclarestmtContext)_localctx).type.t, ((DeclarestmtContext)_localctx).expr.e, false)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(146);
				((DeclarestmtContext)_localctx).VAR = match(VAR);
				setState(147);
				((DeclarestmtContext)_localctx).ID = match(ID);
				setState(148);
				match(IG);
				setState(149);
				((DeclarestmtContext)_localctx).expr = expr(0);
				setState(151);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(150);
					match(PTCOMA);
					}
				}

				_localctx.dec = instructions.NewTodeclare((((DeclarestmtContext)_localctx).VAR!=null?((DeclarestmtContext)_localctx).VAR.getLine():0), (((DeclarestmtContext)_localctx).VAR!=null?((DeclarestmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarestmtContext)_localctx).ID!=null?((DeclarestmtContext)_localctx).ID.getText():null), environment.NULL, ((DeclarestmtContext)_localctx).expr.e, false)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(155);
				((DeclarestmtContext)_localctx).VAR = match(VAR);
				setState(156);
				((DeclarestmtContext)_localctx).ID = match(ID);
				setState(157);
				match(DOUBLEPTS);
				setState(158);
				((DeclarestmtContext)_localctx).type = type();
				setState(159);
				match(QUESTION);
				setState(161);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(160);
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
			setState(187);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(167);
				((ConstantstmtContext)_localctx).LET = match(LET);
				setState(168);
				((ConstantstmtContext)_localctx).ID = match(ID);
				setState(169);
				match(DOUBLEPTS);
				setState(170);
				((ConstantstmtContext)_localctx).type = type();
				setState(171);
				match(IG);
				setState(172);
				((ConstantstmtContext)_localctx).expr = expr(0);
				setState(174);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(173);
					match(PTCOMA);
					}
				}

				_localctx.const = instructions.NewTodeclare((((ConstantstmtContext)_localctx).LET!=null?((ConstantstmtContext)_localctx).LET.getLine():0), (((ConstantstmtContext)_localctx).LET!=null?((ConstantstmtContext)_localctx).LET.getCharPositionInLine():0), (((ConstantstmtContext)_localctx).ID!=null?((ConstantstmtContext)_localctx).ID.getText():null), ((ConstantstmtContext)_localctx).type.t, ((ConstantstmtContext)_localctx).expr.e, true)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(178);
				((ConstantstmtContext)_localctx).LET = match(LET);
				setState(179);
				((ConstantstmtContext)_localctx).ID = match(ID);
				setState(180);
				match(IG);
				setState(181);
				((ConstantstmtContext)_localctx).expr = expr(0);
				setState(183);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(182);
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
			setState(190); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(189);
					((BlockelifsContext)_localctx).ifstmt = ifstmt();
					((BlockelifsContext)_localctx).elif.add(((BlockelifsContext)_localctx).ifstmt);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(192); 
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
			setState(223);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(196);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(197);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(198);
				match(LLAVEIZQ);
				setState(199);
				((IfstmtContext)_localctx).ifb = block();
				setState(200);
				match(LLAVEDER);
				 ((IfstmtContext)_localctx).ift =  instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).ifb.blk, nil); 
				                                        _localctx.el = ((IfstmtContext)_localctx).ifb.blk
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(203);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(204);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(205);
				match(LLAVEIZQ);
				setState(206);
				((IfstmtContext)_localctx).ifelseblck = block();
				setState(207);
				match(LLAVEDER);
				setState(208);
				match(ELSE);
				setState(209);
				match(LLAVEIZQ);
				setState(210);
				((IfstmtContext)_localctx).elseifblck = block();
				setState(211);
				match(LLAVEDER);
				 ((IfstmtContext)_localctx).ift =  instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).ifelseblck.blk, ((IfstmtContext)_localctx).elseifblck.blk); 
				                                                                                        ((IfstmtContext)_localctx).el =  ((IfstmtContext)_localctx).ifelseblck.blk;
				                                                                                        
				                                                                                        
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(214);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(215);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(216);
				match(LLAVEIZQ);
				setState(217);
				((IfstmtContext)_localctx).elif = block();
				setState(218);
				match(LLAVEDER);
				setState(219);
				match(ELSE);
				setState(220);
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
			setState(241);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(225);
				((SwitchstmtContext)_localctx).SWITCH = match(SWITCH);
				setState(226);
				((SwitchstmtContext)_localctx).expr = expr(0);
				setState(227);
				match(LLAVEIZQ);
				setState(228);
				match(DEFAULT);
				setState(229);
				match(DOUBLEPTS);
				setState(230);
				((SwitchstmtContext)_localctx).block = block();
				setState(231);
				match(LLAVEDER);
				 ((SwitchstmtContext)_localctx).switch =  instructions.NewSwitch((((SwitchstmtContext)_localctx).SWITCH!=null?((SwitchstmtContext)_localctx).SWITCH.getLine():0), (((SwitchstmtContext)_localctx).SWITCH!=null?((SwitchstmtContext)_localctx).SWITCH.getCharPositionInLine():0), ((SwitchstmtContext)_localctx).expr.e, nil,  nil, ((SwitchstmtContext)_localctx).block.blk);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(234);
				((SwitchstmtContext)_localctx).SWITCH = match(SWITCH);
				setState(235);
				((SwitchstmtContext)_localctx).expr = expr(0);
				setState(236);
				match(LLAVEIZQ);
				setState(237);
				((SwitchstmtContext)_localctx).casestmt = casestmt();
				setState(238);
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
			setState(259);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(243);
				((CasestmtContext)_localctx).CASE = match(CASE);
				setState(244);
				((CasestmtContext)_localctx).expr = expr(0);
				setState(245);
				match(DOUBLEPTS);
				setState(246);
				((CasestmtContext)_localctx).block = block();
				setState(247);
				((CasestmtContext)_localctx).blockcases = blockcases();
				 _localctx.cas = instructions.NewCase((((CasestmtContext)_localctx).CASE!=null?((CasestmtContext)_localctx).CASE.getLine():0),(((CasestmtContext)_localctx).CASE!=null?((CasestmtContext)_localctx).CASE.getCharPositionInLine():0), ((CasestmtContext)_localctx).expr.e,((CasestmtContext)_localctx).block.blk, ((CasestmtContext)_localctx).blockcases.blkcase)
				                                fmt.Println("entro")
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(250);
				((CasestmtContext)_localctx).CASE = match(CASE);
				setState(251);
				((CasestmtContext)_localctx).expr = expr(0);
				setState(252);
				match(DOUBLEPTS);
				setState(253);
				((CasestmtContext)_localctx).sen = block();
				setState(254);
				match(DEFAULT);
				setState(255);
				match(DOUBLEPTS);
				setState(256);
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
			setState(262); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(261);
					((BlockcasesContext)_localctx).casestmt = casestmt();
					((BlockcasesContext)_localctx).casedef.add(((BlockcasesContext)_localctx).casestmt);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(264); 
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
			setState(268);
			((AssignationstmtContext)_localctx).ID = match(ID);
			setState(269);
			((AssignationstmtContext)_localctx).IG = match(IG);
			setState(270);
			((AssignationstmtContext)_localctx).expr = expr(0);
			setState(272);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PTCOMA) {
				{
				setState(271);
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
			setState(276);
			((WhilestmtContext)_localctx).WHILE = match(WHILE);
			setState(277);
			((WhilestmtContext)_localctx).expr = expr(0);
			setState(278);
			match(LLAVEIZQ);
			setState(279);
			((WhilestmtContext)_localctx).block = block();
			setState(280);
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
			setState(304);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(283);
				((ForstmtContext)_localctx).FOR = match(FOR);
				setState(284);
				((ForstmtContext)_localctx).ID = match(ID);
				setState(285);
				match(IN);
				setState(286);
				((ForstmtContext)_localctx).expr = expr(0);
				setState(287);
				match(LLAVEIZQ);
				setState(288);
				((ForstmtContext)_localctx).block = block();
				setState(289);
				match(LLAVEDER);
				_localctx.for = instructions.NewFor((((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getLine():0), (((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getCharPositionInLine():0), (((ForstmtContext)_localctx).ID!=null?((ForstmtContext)_localctx).ID.getText():null), ((ForstmtContext)_localctx).expr.e, nil, ((ForstmtContext)_localctx).block.blk)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(292);
				((ForstmtContext)_localctx).FOR = match(FOR);
				setState(293);
				((ForstmtContext)_localctx).ID = match(ID);
				setState(294);
				match(IN);
				setState(295);
				((ForstmtContext)_localctx).first = expr(0);
				setState(296);
				match(POINT);
				setState(297);
				match(POINT);
				setState(298);
				((ForstmtContext)_localctx).second = expr(0);
				setState(299);
				match(LLAVEIZQ);
				setState(300);
				((ForstmtContext)_localctx).block = block();
				setState(301);
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
			setState(306);
			((GuardstmtContext)_localctx).GUARD = match(GUARD);
			setState(307);
			((GuardstmtContext)_localctx).expr = expr(0);
			setState(308);
			match(ELSE);
			setState(309);
			match(LLAVEIZQ);
			setState(310);
			((GuardstmtContext)_localctx).block = block();
			setState(311);
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
			setState(312);
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
			setState(337);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,21,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(315);
				((TransferstmtContext)_localctx).RETURN = match(RETURN);
				setState(317);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(316);
					match(PTCOMA);
					}
				}

				_localctx.tran = instructions.NewReturnIn((((TransferstmtContext)_localctx).RETURN!=null?((TransferstmtContext)_localctx).RETURN.getLine():0), (((TransferstmtContext)_localctx).RETURN!=null?((TransferstmtContext)_localctx).RETURN.getCharPositionInLine():0), nil)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(320);
				((TransferstmtContext)_localctx).RETURN = match(RETURN);
				setState(321);
				((TransferstmtContext)_localctx).expr = expr(0);
				setState(323);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(322);
					match(PTCOMA);
					}
				}

				_localctx.tran = instructions.NewReturnIn((((TransferstmtContext)_localctx).RETURN!=null?((TransferstmtContext)_localctx).RETURN.getLine():0), (((TransferstmtContext)_localctx).RETURN!=null?((TransferstmtContext)_localctx).RETURN.getCharPositionInLine():0), ((TransferstmtContext)_localctx).expr.e)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(327);
				((TransferstmtContext)_localctx).CONTINUE = match(CONTINUE);
				setState(329);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(328);
					match(PTCOMA);
					}
				}

				_localctx.tran = instructions.NewContinue((((TransferstmtContext)_localctx).CONTINUE!=null?((TransferstmtContext)_localctx).CONTINUE.getLine():0), (((TransferstmtContext)_localctx).CONTINUE!=null?((TransferstmtContext)_localctx).CONTINUE.getCharPositionInLine():0))
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(332);
				((TransferstmtContext)_localctx).BREAK = match(BREAK);
				setState(334);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(333);
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
			setState(392);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,26,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(339);
				((DeclarevectorstmtContext)_localctx).VAR = match(VAR);
				setState(340);
				((DeclarevectorstmtContext)_localctx).ID = match(ID);
				setState(341);
				match(DOUBLEPTS);
				setState(342);
				match(CORCHETEIZQ);
				setState(343);
				((DeclarevectorstmtContext)_localctx).type = type();
				setState(344);
				match(CORCHETEDER);
				setState(345);
				match(IG);
				setState(346);
				match(CORCHETEIZQ);
				setState(347);
				((DeclarevectorstmtContext)_localctx).listParams = listParams(0);
				setState(348);
				match(CORCHETEDER);
				setState(350);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(349);
					match(PTCOMA);
					}
				}

				_localctx.decvec = instructions.NewToDeclareVector((((DeclarevectorstmtContext)_localctx).VAR!=null?((DeclarevectorstmtContext)_localctx).VAR.getLine():0), (((DeclarevectorstmtContext)_localctx).VAR!=null?((DeclarevectorstmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarevectorstmtContext)_localctx).ID!=null?((DeclarevectorstmtContext)_localctx).ID.getText():null), ((DeclarevectorstmtContext)_localctx).type.t, ((DeclarevectorstmtContext)_localctx).listParams.l, "")
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(354);
				((DeclarevectorstmtContext)_localctx).VAR = match(VAR);
				setState(355);
				((DeclarevectorstmtContext)_localctx).ID = match(ID);
				setState(356);
				match(DOUBLEPTS);
				setState(357);
				match(CORCHETEIZQ);
				setState(358);
				((DeclarevectorstmtContext)_localctx).type = type();
				setState(359);
				match(CORCHETEDER);
				setState(360);
				match(IG);
				setState(361);
				match(CORCHETEIZQ);
				setState(362);
				match(CORCHETEDER);
				setState(364);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(363);
					match(PTCOMA);
					}
				}

				_localctx.decvec = instructions.NewToDeclareVector((((DeclarevectorstmtContext)_localctx).VAR!=null?((DeclarevectorstmtContext)_localctx).VAR.getLine():0), (((DeclarevectorstmtContext)_localctx).VAR!=null?((DeclarevectorstmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarevectorstmtContext)_localctx).ID!=null?((DeclarevectorstmtContext)_localctx).ID.getText():null), ((DeclarevectorstmtContext)_localctx).type.t, nil, "")
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(368);
				match(VAR);
				setState(369);
				match(ID);
				setState(370);
				match(IG);
				setState(371);
				match(CORCHETEIZQ);
				setState(372);
				type();
				setState(373);
				match(CORCHETEDER);
				setState(374);
				match(PARIZQ);
				setState(375);
				match(PARDER);
				setState(377);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(376);
					match(PTCOMA);
					}
				}

				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(379);
				((DeclarevectorstmtContext)_localctx).VAR = match(VAR);
				setState(380);
				((DeclarevectorstmtContext)_localctx).F = match(ID);
				setState(381);
				match(DOUBLEPTS);
				setState(382);
				match(CORCHETEIZQ);
				setState(383);
				((DeclarevectorstmtContext)_localctx).type = type();
				setState(384);
				match(CORCHETEDER);
				setState(385);
				match(IG);
				setState(386);
				((DeclarevectorstmtContext)_localctx).S = match(ID);
				setState(388);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(387);
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
		enterRule(_localctx, 34, RULE_appendstmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(394);
			((AppendstmtContext)_localctx).ID = match(ID);
			setState(395);
			match(POINT);
			setState(396);
			match(APPEND);
			setState(397);
			match(PARIZQ);
			setState(398);
			((AppendstmtContext)_localctx).expr = expr(0);
			setState(399);
			match(PARDER);
			setState(401);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PTCOMA) {
				{
				setState(400);
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
		enterRule(_localctx, 36, RULE_removelaststmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(405);
			((RemovelaststmtContext)_localctx).ID = match(ID);
			setState(406);
			match(POINT);
			setState(407);
			match(REMOVELAST);
			setState(408);
			match(PARIZQ);
			setState(409);
			match(PARDER);
			setState(411);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PTCOMA) {
				{
				setState(410);
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
		enterRule(_localctx, 38, RULE_removestmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(415);
			((RemovestmtContext)_localctx).ID = match(ID);
			setState(416);
			match(POINT);
			setState(417);
			match(REMOVE);
			setState(418);
			match(PARIZQ);
			setState(419);
			match(AT);
			setState(420);
			match(DOUBLEPTS);
			setState(421);
			((RemovestmtContext)_localctx).expr = expr(0);
			setState(422);
			match(PARDER);
			setState(424);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PTCOMA) {
				{
				setState(423);
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
		enterRule(_localctx, 40, RULE_emptvecstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(428);
			((EmptvecstmtContext)_localctx).ID = match(ID);
			setState(429);
			match(POINT);
			setState(430);
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
		enterRule(_localctx, 42, RULE_countvecstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(433);
			((CountvecstmtContext)_localctx).ID = match(ID);
			setState(434);
			match(POINT);
			setState(435);
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
		enterRule(_localctx, 44, RULE_accessvecstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(438);
			((AccessvecstmtContext)_localctx).ID = match(ID);
			setState(439);
			match(CORCHETEIZQ);
			setState(440);
			((AccessvecstmtContext)_localctx).expr = expr(0);
			setState(441);
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
		enterRule(_localctx, 46, RULE_assignationvecstmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(444);
			((AssignationvecstmtContext)_localctx).ID = match(ID);
			setState(445);
			match(CORCHETEIZQ);
			setState(446);
			((AssignationvecstmtContext)_localctx).expprim = expr(0);
			setState(447);
			match(CORCHETEDER);
			setState(448);
			match(IG);
			setState(449);
			((AssignationvecstmtContext)_localctx).expsegundo = expr(0);
			setState(451);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PTCOMA) {
				{
				setState(450);
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
		enterRule(_localctx, 48, RULE_declarematrixstmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(455);
			match(VAR);
			setState(456);
			match(ID);
			setState(457);
			match(PARIZQ);
			setState(458);
			match(DOUBLEPTS);
			setState(459);
			type();
			setState(460);
			match(PARDER);
			setState(461);
			match(IG);
			setState(463);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PTCOMA) {
				{
				setState(462);
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
		enterRule(_localctx, 50, RULE_declarefuncstmt);
		try {
			setState(509);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,32,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(465);
				match(FUNC);
				setState(466);
				((DeclarefuncstmtContext)_localctx).ID = match(ID);
				setState(467);
				match(PARIZQ);
				setState(468);
				((DeclarefuncstmtContext)_localctx).listParams = listParams(0);
				setState(469);
				match(PARDER);
				setState(470);
				match(SUB);
				setState(471);
				match(MAYOR);
				setState(472);
				((DeclarefuncstmtContext)_localctx).type = type();
				setState(473);
				match(LLAVEIZQ);
				setState(474);
				((DeclarefuncstmtContext)_localctx).block = block();
				setState(475);
				match(LLAVEDER);
				_localctx.decfunc = instructions.NewToDeclareFunction((((DeclarefuncstmtContext)_localctx).ID!=null?((DeclarefuncstmtContext)_localctx).ID.getLine():0), (((DeclarefuncstmtContext)_localctx).ID!=null?((DeclarefuncstmtContext)_localctx).ID.getCharPositionInLine():0), (((DeclarefuncstmtContext)_localctx).ID!=null?((DeclarefuncstmtContext)_localctx).ID.getText():null), ((DeclarefuncstmtContext)_localctx).listParams.l, ((DeclarefuncstmtContext)_localctx).type.t, ((DeclarefuncstmtContext)_localctx).block.blk, 1)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(478);
				match(FUNC);
				setState(479);
				((DeclarefuncstmtContext)_localctx).ID = match(ID);
				setState(480);
				match(PARIZQ);
				setState(481);
				((DeclarefuncstmtContext)_localctx).listParams = listParams(0);
				setState(482);
				match(PARDER);
				setState(483);
				match(LLAVEIZQ);
				setState(484);
				((DeclarefuncstmtContext)_localctx).block = block();
				setState(485);
				match(LLAVEDER);
				_localctx.decfunc = instructions.NewToDeclareFunction((((DeclarefuncstmtContext)_localctx).ID!=null?((DeclarefuncstmtContext)_localctx).ID.getLine():0), (((DeclarefuncstmtContext)_localctx).ID!=null?((DeclarefuncstmtContext)_localctx).ID.getCharPositionInLine():0), (((DeclarefuncstmtContext)_localctx).ID!=null?((DeclarefuncstmtContext)_localctx).ID.getText():null), ((DeclarefuncstmtContext)_localctx).listParams.l, environment.NULL, ((DeclarefuncstmtContext)_localctx).block.blk, 2)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(488);
				match(FUNC);
				setState(489);
				((DeclarefuncstmtContext)_localctx).ID = match(ID);
				setState(490);
				match(PARIZQ);
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
				_localctx.decfunc = instructions.NewToDeclareFunction((((DeclarefuncstmtContext)_localctx).ID!=null?((DeclarefuncstmtContext)_localctx).ID.getLine():0), (((DeclarefuncstmtContext)_localctx).ID!=null?((DeclarefuncstmtContext)_localctx).ID.getCharPositionInLine():0), (((DeclarefuncstmtContext)_localctx).ID!=null?((DeclarefuncstmtContext)_localctx).ID.getText():null), nil, ((DeclarefuncstmtContext)_localctx).type.t, ((DeclarefuncstmtContext)_localctx).block.blk, 3)
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(500);
				match(FUNC);
				setState(501);
				((DeclarefuncstmtContext)_localctx).ID = match(ID);
				setState(502);
				match(PARIZQ);
				setState(503);
				match(PARDER);
				setState(504);
				match(LLAVEIZQ);
				setState(505);
				((DeclarefuncstmtContext)_localctx).block = block();
				setState(506);
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
		int _startState = 52;
		enterRecursionRule(_localctx, 52, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(547);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,33,_ctx) ) {
			case 1:
				{
				setState(512);
				((ExprContext)_localctx).op = match(NOT);
				setState(513);
				((ExprContext)_localctx).left = ((ExprContext)_localctx).expr = expr(12);
				 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).left.e) 
				}
				break;
			case 2:
				{
				setState(516);
				((ExprContext)_localctx).op = match(SUB);
				setState(517);
				((ExprContext)_localctx).left = ((ExprContext)_localctx).expr = expr(11);
				 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e,  (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), nil)
				}
				break;
			case 3:
				{
				setState(520);
				match(PARIZQ);
				setState(521);
				((ExprContext)_localctx).expr = expr(0);
				setState(522);
				match(PARDER);
				 _localctx.e = ((ExprContext)_localctx).expr.e 
				}
				break;
			case 4:
				{
				setState(525);
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
				setState(527);
				((ExprContext)_localctx).STRING = match(STRING);

				        str := (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getText():null)
				        _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getLine():0), (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getCharPositionInLine():0), str[1:len(str)-1],environment.STRING)
				    
				}
				break;
			case 6:
				{
				setState(529);
				((ExprContext)_localctx).TRUE = match(TRUE);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).TRUE!=null?((ExprContext)_localctx).TRUE.getLine():0), (((ExprContext)_localctx).TRUE!=null?((ExprContext)_localctx).TRUE.getCharPositionInLine():0), true, environment.BOOLEAN) 
				}
				break;
			case 7:
				{
				setState(531);
				((ExprContext)_localctx).FALSE = match(FALSE);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).FALSE!=null?((ExprContext)_localctx).FALSE.getLine():0), (((ExprContext)_localctx).FALSE!=null?((ExprContext)_localctx).FALSE.getCharPositionInLine():0), false, environment.BOOLEAN) 
				}
				break;
			case 8:
				{
				setState(533);
				((ExprContext)_localctx).NIL = match(NIL);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).NIL!=null?((ExprContext)_localctx).NIL.getLine():0), (((ExprContext)_localctx).NIL!=null?((ExprContext)_localctx).NIL.getCharPositionInLine():0), nil, environment.NULL) 
				}
				break;
			case 9:
				{
				setState(535);
				((ExprContext)_localctx).accessstmt = accessstmt();
				_localctx.e = ((ExprContext)_localctx).accessstmt.access
				}
				break;
			case 10:
				{
				setState(538);
				((ExprContext)_localctx).emptvecstmt = emptvecstmt();
				_localctx.e = ((ExprContext)_localctx).emptvecstmt.emptyvec
				}
				break;
			case 11:
				{
				setState(541);
				((ExprContext)_localctx).countvecstmt = countvecstmt();
				_localctx.e = ((ExprContext)_localctx).countvecstmt.count
				}
				break;
			case 12:
				{
				setState(544);
				((ExprContext)_localctx).accessvecstmt = accessvecstmt();
				_localctx.e = ((ExprContext)_localctx).accessvecstmt.accessvec
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(586);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,35,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(584);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,34,_ctx) ) {
					case 1:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(549);
						if (!(precpred(_ctx, 19))) throw new FailedPredicateException(this, "precpred(_ctx, 19)");
						setState(550);
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
						setState(551);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(20);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 2:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(554);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(555);
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
						setState(556);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(19);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 3:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(559);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(560);
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
						setState(561);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(18);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 4:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(564);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(565);
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
						setState(566);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(17);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 5:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(569);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(570);
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
						setState(571);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(16);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 6:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(574);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(575);
						((ExprContext)_localctx).op = match(AND);
						setState(576);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(15);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					case 7:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(579);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(580);
						((ExprContext)_localctx).op = match(OR);
						setState(581);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(14);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					}
					} 
				}
				setState(588);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,35,_ctx);
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
		enterRule(_localctx, 54, RULE_accessstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(589);
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
		enterRule(_localctx, 56, RULE_increaseanddecreasestmt);
		int _la;
		try {
			setState(608);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,38,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(592);
				((IncreaseanddecreasestmtContext)_localctx).ID = match(ID);
				setState(593);
				((IncreaseanddecreasestmtContext)_localctx).IG_ADD = match(IG_ADD);
				setState(594);
				((IncreaseanddecreasestmtContext)_localctx).expr = expr(0);
				setState(596);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(595);
					match(PTCOMA);
					}
				}

				_localctx.increasedecrease = instructions.NewIncreaseDecrease((((IncreaseanddecreasestmtContext)_localctx).ID!=null?((IncreaseanddecreasestmtContext)_localctx).ID.getLine():0), (((IncreaseanddecreasestmtContext)_localctx).ID!=null?((IncreaseanddecreasestmtContext)_localctx).ID.getCharPositionInLine():0), (((IncreaseanddecreasestmtContext)_localctx).ID!=null?((IncreaseanddecreasestmtContext)_localctx).ID.getText():null), (((IncreaseanddecreasestmtContext)_localctx).IG_ADD!=null?((IncreaseanddecreasestmtContext)_localctx).IG_ADD.getText():null), ((IncreaseanddecreasestmtContext)_localctx).expr.e)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(600);
				((IncreaseanddecreasestmtContext)_localctx).ID = match(ID);
				setState(601);
				((IncreaseanddecreasestmtContext)_localctx).IG_SUB = match(IG_SUB);
				setState(602);
				((IncreaseanddecreasestmtContext)_localctx).expr = expr(0);
				setState(604);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(603);
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
		enterRule(_localctx, 58, RULE_type);
		try {
			setState(620);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRINGS:
				enterOuterAlt(_localctx, 1);
				{
				setState(610);
				match(STRINGS);
				_localctx.t = environment.STRING
				}
				break;
			case INTS:
				enterOuterAlt(_localctx, 2);
				{
				setState(612);
				match(INTS);
				_localctx.t = environment.INTEGER
				}
				break;
			case FLOATS:
				enterOuterAlt(_localctx, 3);
				{
				setState(614);
				match(FLOATS);
				_localctx.t = environment.FLOAT
				}
				break;
			case BOOLS:
				enterOuterAlt(_localctx, 4);
				{
				setState(616);
				match(BOOLS);
				_localctx.t = environment.BOOLEAN 
				}
				break;
			case CHARACTERS:
				enterOuterAlt(_localctx, 5);
				{
				setState(618);
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
		int _startState = 60;
		enterRecursionRule(_localctx, 60, RULE_listParams, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(623);
			((ListParamsContext)_localctx).expr = expr(0);

			            _localctx.l = []interface{}{}
			            _localctx.l = append(_localctx.l, ((ListParamsContext)_localctx).expr.e)
			        
			}
			_ctx.stop = _input.LT(-1);
			setState(633);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,40,_ctx);
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
					setState(626);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(627);
					match(COMA);
					setState(628);
					((ListParamsContext)_localctx).expr = expr(0);

					                                          var arr []interface{}
					                                          arr = append(((ListParamsContext)_localctx).list.l, ((ListParamsContext)_localctx).expr.e)
					                                          _localctx.l = arr
					                                      
					}
					} 
				}
				setState(635);
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

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 26:
			return expr_sempred((ExprContext)_localctx, predIndex);
		case 30:
			return listParams_sempred((ListParamsContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 19);
		case 1:
			return precpred(_ctx, 18);
		case 2:
			return precpred(_ctx, 17);
		case 3:
			return precpred(_ctx, 16);
		case 4:
			return precpred(_ctx, 15);
		case 5:
			return precpred(_ctx, 14);
		case 6:
			return precpred(_ctx, 13);
		}
		return true;
	}
	private boolean listParams_sempred(ListParamsContext _localctx, int predIndex) {
		switch (predIndex) {
		case 7:
			return precpred(_ctx, 2);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3C\u027f\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \3\2"+
		"\3\2\3\2\3\2\3\3\6\3F\n\3\r\3\16\3G\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4\177\n\4\3\5\3\5\3\5\3\5\3\5"+
		"\5\5\u0086\n\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6\u0091\n\6\3\6\3"+
		"\6\3\6\3\6\3\6\3\6\3\6\5\6\u009a\n\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5"+
		"\6\u00a4\n\6\3\6\3\6\5\6\u00a8\n\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\5\7\u00b1"+
		"\n\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\5\7\u00ba\n\7\3\7\3\7\5\7\u00be\n\7\3"+
		"\b\6\b\u00c1\n\b\r\b\16\b\u00c2\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t"+
		"\3\t\3\t\5\t\u00e2\n\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n"+
		"\3\n\3\n\3\n\3\n\5\n\u00f4\n\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\5\13\u0106\n\13\3\f\6\f\u0109"+
		"\n\f\r\f\16\f\u010a\3\f\3\f\3\r\3\r\3\r\3\r\5\r\u0113\n\r\3\r\3\r\3\16"+
		"\3\16\3\16\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17"+
		"\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\5\17"+
		"\u0133\n\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\21\3\21\5\21"+
		"\u0140\n\21\3\21\3\21\3\21\3\21\5\21\u0146\n\21\3\21\3\21\3\21\3\21\5"+
		"\21\u014c\n\21\3\21\3\21\3\21\5\21\u0151\n\21\3\21\5\21\u0154\n\21\3\22"+
		"\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\5\22\u0161\n\22\3\22"+
		"\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\5\22\u016f\n\22"+
		"\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\5\22\u017c\n\22"+
		"\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\5\22\u0187\n\22\3\22\3\22"+
		"\5\22\u018b\n\22\3\23\3\23\3\23\3\23\3\23\3\23\3\23\5\23\u0194\n\23\3"+
		"\23\3\23\3\24\3\24\3\24\3\24\3\24\3\24\5\24\u019e\n\24\3\24\3\24\3\25"+
		"\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\5\25\u01ab\n\25\3\25\3\25\3\26"+
		"\3\26\3\26\3\26\3\26\3\27\3\27\3\27\3\27\3\27\3\30\3\30\3\30\3\30\3\30"+
		"\3\30\3\31\3\31\3\31\3\31\3\31\3\31\3\31\5\31\u01c6\n\31\3\31\3\31\3\32"+
		"\3\32\3\32\3\32\3\32\3\32\3\32\3\32\5\32\u01d2\n\32\3\33\3\33\3\33\3\33"+
		"\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33"+
		"\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33"+
		"\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33\5\33\u0200"+
		"\n\33\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34"+
		"\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34"+
		"\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\5\34\u0226\n\34\3\34\3\34"+
		"\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34"+
		"\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34"+
		"\3\34\3\34\3\34\3\34\3\34\7\34\u024b\n\34\f\34\16\34\u024e\13\34\3\35"+
		"\3\35\3\35\3\36\3\36\3\36\3\36\5\36\u0257\n\36\3\36\3\36\3\36\3\36\3\36"+
		"\3\36\5\36\u025f\n\36\3\36\3\36\5\36\u0263\n\36\3\37\3\37\3\37\3\37\3"+
		"\37\3\37\3\37\3\37\3\37\3\37\5\37\u026f\n\37\3 \3 \3 \3 \3 \3 \3 \3 \3"+
		" \7 \u027a\n \f \16 \u027d\13 \3 \2\4\66>!\2\4\6\b\n\f\16\20\22\24\26"+
		"\30\32\34\36 \"$&(*,.\60\62\64\668:<>\2\b\3\2\27\31\3\2\61\62\3\2\63\65"+
		"\4\2--//\4\2..\60\60\3\2%&\2\u02b1\2@\3\2\2\2\4E\3\2\2\2\6~\3\2\2\2\b"+
		"\u0080\3\2\2\2\n\u00a7\3\2\2\2\f\u00bd\3\2\2\2\16\u00c0\3\2\2\2\20\u00e1"+
		"\3\2\2\2\22\u00f3\3\2\2\2\24\u0105\3\2\2\2\26\u0108\3\2\2\2\30\u010e\3"+
		"\2\2\2\32\u0116\3\2\2\2\34\u0132\3\2\2\2\36\u0134\3\2\2\2 \u0153\3\2\2"+
		"\2\"\u018a\3\2\2\2$\u018c\3\2\2\2&\u0197\3\2\2\2(\u01a1\3\2\2\2*\u01ae"+
		"\3\2\2\2,\u01b3\3\2\2\2.\u01b8\3\2\2\2\60\u01be\3\2\2\2\62\u01c9\3\2\2"+
		"\2\64\u01ff\3\2\2\2\66\u0225\3\2\2\28\u024f\3\2\2\2:\u0262\3\2\2\2<\u026e"+
		"\3\2\2\2>\u0270\3\2\2\2@A\5\4\3\2AB\7\2\2\3BC\b\2\1\2C\3\3\2\2\2DF\5\6"+
		"\4\2ED\3\2\2\2FG\3\2\2\2GE\3\2\2\2GH\3\2\2\2HI\3\2\2\2IJ\b\3\1\2J\5\3"+
		"\2\2\2KL\5\b\5\2LM\b\4\1\2M\177\3\2\2\2NO\5\n\6\2OP\b\4\1\2P\177\3\2\2"+
		"\2QR\5\f\7\2RS\b\4\1\2S\177\3\2\2\2TU\5\30\r\2UV\b\4\1\2V\177\3\2\2\2"+
		"WX\5:\36\2XY\b\4\1\2Y\177\3\2\2\2Z[\5\20\t\2[\\\b\4\1\2\\\177\3\2\2\2"+
		"]^\5\32\16\2^_\b\4\1\2_\177\3\2\2\2`a\5\34\17\2ab\b\4\1\2b\177\3\2\2\2"+
		"cd\5\22\n\2de\b\4\1\2e\177\3\2\2\2fg\5\36\20\2gh\b\4\1\2h\177\3\2\2\2"+
		"ij\5 \21\2jk\b\4\1\2k\177\3\2\2\2lm\5\"\22\2mn\b\4\1\2n\177\3\2\2\2op"+
		"\5$\23\2pq\b\4\1\2q\177\3\2\2\2rs\5&\24\2st\b\4\1\2t\177\3\2\2\2uv\5("+
		"\25\2vw\b\4\1\2w\177\3\2\2\2xy\5\60\31\2yz\b\4\1\2z\177\3\2\2\2{|\5\64"+
		"\33\2|}\b\4\1\2}\177\3\2\2\2~K\3\2\2\2~N\3\2\2\2~Q\3\2\2\2~T\3\2\2\2~"+
		"W\3\2\2\2~Z\3\2\2\2~]\3\2\2\2~`\3\2\2\2~c\3\2\2\2~f\3\2\2\2~i\3\2\2\2"+
		"~l\3\2\2\2~o\3\2\2\2~r\3\2\2\2~u\3\2\2\2~x\3\2\2\2~{\3\2\2\2\177\7\3\2"+
		"\2\2\u0080\u0081\7\f\2\2\u0081\u0082\7\66\2\2\u0082\u0083\5\66\34\2\u0083"+
		"\u0085\7\67\2\2\u0084\u0086\7>\2\2\u0085\u0084\3\2\2\2\u0085\u0086\3\2"+
		"\2\2\u0086\u0087\3\2\2\2\u0087\u0088\b\5\1\2\u0088\t\3\2\2\2\u0089\u008a"+
		"\7\b\2\2\u008a\u008b\7$\2\2\u008b\u008c\7:\2\2\u008c\u008d\5<\37\2\u008d"+
		"\u008e\7*\2\2\u008e\u0090\5\66\34\2\u008f\u0091\7>\2\2\u0090\u008f\3\2"+
		"\2\2\u0090\u0091\3\2\2\2\u0091\u0092\3\2\2\2\u0092\u0093\b\6\1\2\u0093"+
		"\u00a8\3\2\2\2\u0094\u0095\7\b\2\2\u0095\u0096\7$\2\2\u0096\u0097\7*\2"+
		"\2\u0097\u0099\5\66\34\2\u0098\u009a\7>\2\2\u0099\u0098\3\2\2\2\u0099"+
		"\u009a\3\2\2\2\u009a\u009b\3\2\2\2\u009b\u009c\b\6\1\2\u009c\u00a8\3\2"+
		"\2\2\u009d\u009e\7\b\2\2\u009e\u009f\7$\2\2\u009f\u00a0\7:\2\2\u00a0\u00a1"+
		"\5<\37\2\u00a1\u00a3\7;\2\2\u00a2\u00a4\7>\2\2\u00a3\u00a2\3\2\2\2\u00a3"+
		"\u00a4\3\2\2\2\u00a4\u00a5\3\2\2\2\u00a5\u00a6\b\6\1\2\u00a6\u00a8\3\2"+
		"\2\2\u00a7\u0089\3\2\2\2\u00a7\u0094\3\2\2\2\u00a7\u009d\3\2\2\2\u00a8"+
		"\13\3\2\2\2\u00a9\u00aa\7\20\2\2\u00aa\u00ab\7$\2\2\u00ab\u00ac\7:\2\2"+
		"\u00ac\u00ad\5<\37\2\u00ad\u00ae\7*\2\2\u00ae\u00b0\5\66\34\2\u00af\u00b1"+
		"\7>\2\2\u00b0\u00af\3\2\2\2\u00b0\u00b1\3\2\2\2\u00b1\u00b2\3\2\2\2\u00b2"+
		"\u00b3\b\7\1\2\u00b3\u00be\3\2\2\2\u00b4\u00b5\7\20\2\2\u00b5\u00b6\7"+
		"$\2\2\u00b6\u00b7\7*\2\2\u00b7\u00b9\5\66\34\2\u00b8\u00ba\7>\2\2\u00b9"+
		"\u00b8\3\2\2\2\u00b9\u00ba\3\2\2\2\u00ba\u00bb\3\2\2\2\u00bb\u00bc\b\7"+
		"\1\2\u00bc\u00be\3\2\2\2\u00bd\u00a9\3\2\2\2\u00bd\u00b4\3\2\2\2\u00be"+
		"\r\3\2\2\2\u00bf\u00c1\5\20\t\2\u00c0\u00bf\3\2\2\2\u00c1\u00c2\3\2\2"+
		"\2\u00c2\u00c0\3\2\2\2\u00c2\u00c3\3\2\2\2\u00c3\u00c4\3\2\2\2\u00c4\u00c5"+
		"\b\b\1\2\u00c5\17\3\2\2\2\u00c6\u00c7\7\r\2\2\u00c7\u00c8\5\66\34\2\u00c8"+
		"\u00c9\78\2\2\u00c9\u00ca\5\4\3\2\u00ca\u00cb\79\2\2\u00cb\u00cc\b\t\1"+
		"\2\u00cc\u00e2\3\2\2\2\u00cd\u00ce\7\r\2\2\u00ce\u00cf\5\66\34\2\u00cf"+
		"\u00d0\78\2\2\u00d0\u00d1\5\4\3\2\u00d1\u00d2\79\2\2\u00d2\u00d3\7\16"+
		"\2\2\u00d3\u00d4\78\2\2\u00d4\u00d5\5\4\3\2\u00d5\u00d6\79\2\2\u00d6\u00d7"+
		"\b\t\1\2\u00d7\u00e2\3\2\2\2\u00d8\u00d9\7\r\2\2\u00d9\u00da\5\66\34\2"+
		"\u00da\u00db\78\2\2\u00db\u00dc\5\4\3\2\u00dc\u00dd\79\2\2\u00dd\u00de"+
		"\7\16\2\2\u00de\u00df\5\16\b\2\u00df\u00e0\b\t\1\2\u00e0\u00e2\3\2\2\2"+
		"\u00e1\u00c6\3\2\2\2\u00e1\u00cd\3\2\2\2\u00e1\u00d8\3\2\2\2\u00e2\21"+
		"\3\2\2\2\u00e3\u00e4\7\23\2\2\u00e4\u00e5\5\66\34\2\u00e5\u00e6\78\2\2"+
		"\u00e6\u00e7\7\25\2\2\u00e7\u00e8\7:\2\2\u00e8\u00e9\5\4\3\2\u00e9\u00ea"+
		"\79\2\2\u00ea\u00eb\b\n\1\2\u00eb\u00f4\3\2\2\2\u00ec\u00ed\7\23\2\2\u00ed"+
		"\u00ee\5\66\34\2\u00ee\u00ef\78\2\2\u00ef\u00f0\5\24\13\2\u00f0\u00f1"+
		"\79\2\2\u00f1\u00f2\b\n\1\2\u00f2\u00f4\3\2\2\2\u00f3\u00e3\3\2\2\2\u00f3"+
		"\u00ec\3\2\2\2\u00f4\23\3\2\2\2\u00f5\u00f6\7\24\2\2\u00f6\u00f7\5\66"+
		"\34\2\u00f7\u00f8\7:\2\2\u00f8\u00f9\5\4\3\2\u00f9\u00fa\5\26\f\2\u00fa"+
		"\u00fb\b\13\1\2\u00fb\u0106\3\2\2\2\u00fc\u00fd\7\24\2\2\u00fd\u00fe\5"+
		"\66\34\2\u00fe\u00ff\7:\2\2\u00ff\u0100\5\4\3\2\u0100\u0101\7\25\2\2\u0101"+
		"\u0102\7:\2\2\u0102\u0103\5\4\3\2\u0103\u0104\b\13\1\2\u0104\u0106\3\2"+
		"\2\2\u0105\u00f5\3\2\2\2\u0105\u00fc\3\2\2\2\u0106\25\3\2\2\2\u0107\u0109"+
		"\5\24\13\2\u0108\u0107\3\2\2\2\u0109\u010a\3\2\2\2\u010a\u0108\3\2\2\2"+
		"\u010a\u010b\3\2\2\2\u010b\u010c\3\2\2\2\u010c\u010d\b\f\1\2\u010d\27"+
		"\3\2\2\2\u010e\u010f\7$\2\2\u010f\u0110\7*\2\2\u0110\u0112\5\66\34\2\u0111"+
		"\u0113\7>\2\2\u0112\u0111\3\2\2\2\u0112\u0113\3\2\2\2\u0113\u0114\3\2"+
		"\2\2\u0114\u0115\b\r\1\2\u0115\31\3\2\2\2\u0116\u0117\7\17\2\2\u0117\u0118"+
		"\5\66\34\2\u0118\u0119\78\2\2\u0119\u011a\5\4\3\2\u011a\u011b\79\2\2\u011b"+
		"\u011c\b\16\1\2\u011c\33\3\2\2\2\u011d\u011e\7\21\2\2\u011e\u011f\7$\2"+
		"\2\u011f\u0120\7\22\2\2\u0120\u0121\5\66\34\2\u0121\u0122\78\2\2\u0122"+
		"\u0123\5\4\3\2\u0123\u0124\79\2\2\u0124\u0125\b\17\1\2\u0125\u0133\3\2"+
		"\2\2\u0126\u0127\7\21\2\2\u0127\u0128\7$\2\2\u0128\u0129\7\22\2\2\u0129"+
		"\u012a\5\66\34\2\u012a\u012b\7<\2\2\u012b\u012c\7<\2\2\u012c\u012d\5\66"+
		"\34\2\u012d\u012e\78\2\2\u012e\u012f\5\4\3\2\u012f\u0130\79\2\2\u0130"+
		"\u0131\b\17\1\2\u0131\u0133\3\2\2\2\u0132\u011d\3\2\2\2\u0132\u0126\3"+
		"\2\2\2\u0133\35\3\2\2\2\u0134\u0135\7\26\2\2\u0135\u0136\5\66\34\2\u0136"+
		"\u0137\7\16\2\2\u0137\u0138\78\2\2\u0138\u0139\5\4\3\2\u0139\u013a\t\2"+
		"\2\2\u013a\u013b\79\2\2\u013b\u013c\b\20\1\2\u013c\37\3\2\2\2\u013d\u013f"+
		"\7\30\2\2\u013e\u0140\7>\2\2\u013f\u013e\3\2\2\2\u013f\u0140\3\2\2\2\u0140"+
		"\u0141\3\2\2\2\u0141\u0154\b\21\1\2\u0142\u0143\7\30\2\2\u0143\u0145\5"+
		"\66\34\2\u0144\u0146\7>\2\2\u0145\u0144\3\2\2\2\u0145\u0146\3\2\2\2\u0146"+
		"\u0147\3\2\2\2\u0147\u0148\b\21\1\2\u0148\u0154\3\2\2\2\u0149\u014b\7"+
		"\27\2\2\u014a\u014c\7>\2\2\u014b\u014a\3\2\2\2\u014b\u014c\3\2\2\2\u014c"+
		"\u014d\3\2\2\2\u014d\u0154\b\21\1\2\u014e\u0150\7\31\2\2\u014f\u0151\7"+
		">\2\2\u0150\u014f\3\2\2\2\u0150\u0151\3\2\2\2\u0151\u0152\3\2\2\2\u0152"+
		"\u0154\b\21\1\2\u0153\u013d\3\2\2\2\u0153\u0142\3\2\2\2\u0153\u0149\3"+
		"\2\2\2\u0153\u014e\3\2\2\2\u0154!\3\2\2\2\u0155\u0156\7\b\2\2\u0156\u0157"+
		"\7$\2\2\u0157\u0158\7:\2\2\u0158\u0159\7?\2\2\u0159\u015a\5<\37\2\u015a"+
		"\u015b\7@\2\2\u015b\u015c\7*\2\2\u015c\u015d\7?\2\2\u015d\u015e\5> \2"+
		"\u015e\u0160\7@\2\2\u015f\u0161\7>\2\2\u0160\u015f\3\2\2\2\u0160\u0161"+
		"\3\2\2\2\u0161\u0162\3\2\2\2\u0162\u0163\b\22\1\2\u0163\u018b\3\2\2\2"+
		"\u0164\u0165\7\b\2\2\u0165\u0166\7$\2\2\u0166\u0167\7:\2\2\u0167\u0168"+
		"\7?\2\2\u0168\u0169\5<\37\2\u0169\u016a\7@\2\2\u016a\u016b\7*\2\2\u016b"+
		"\u016c\7?\2\2\u016c\u016e\7@\2\2\u016d\u016f\7>\2\2\u016e\u016d\3\2\2"+
		"\2\u016e\u016f\3\2\2\2\u016f\u0170\3\2\2\2\u0170\u0171\b\22\1\2\u0171"+
		"\u018b\3\2\2\2\u0172\u0173\7\b\2\2\u0173\u0174\7$\2\2\u0174\u0175\7*\2"+
		"\2\u0175\u0176\7?\2\2\u0176\u0177\5<\37\2\u0177\u0178\7@\2\2\u0178\u0179"+
		"\7\66\2\2\u0179\u017b\7\67\2\2\u017a\u017c\7>\2\2\u017b\u017a\3\2\2\2"+
		"\u017b\u017c\3\2\2\2\u017c\u018b\3\2\2\2\u017d\u017e\7\b\2\2\u017e\u017f"+
		"\7$\2\2\u017f\u0180\7:\2\2\u0180\u0181\7?\2\2\u0181\u0182\5<\37\2\u0182"+
		"\u0183\7@\2\2\u0183\u0184\7*\2\2\u0184\u0186\7$\2\2\u0185\u0187\7>\2\2"+
		"\u0186\u0185\3\2\2\2\u0186\u0187\3\2\2\2\u0187\u0188\3\2\2\2\u0188\u0189"+
		"\b\22\1\2\u0189\u018b\3\2\2\2\u018a\u0155\3\2\2\2\u018a\u0164\3\2\2\2"+
		"\u018a\u0172\3\2\2\2\u018a\u017d\3\2\2\2\u018b#\3\2\2\2\u018c\u018d\7"+
		"$\2\2\u018d\u018e\7<\2\2\u018e\u018f\7\32\2\2\u018f\u0190\7\66\2\2\u0190"+
		"\u0191\5\66\34\2\u0191\u0193\7\67\2\2\u0192\u0194\7>\2\2\u0193\u0192\3"+
		"\2\2\2\u0193\u0194\3\2\2\2\u0194\u0195\3\2\2\2\u0195\u0196\b\23\1\2\u0196"+
		"%\3\2\2\2\u0197\u0198\7$\2\2\u0198\u0199\7<\2\2\u0199\u019a\7\33\2\2\u019a"+
		"\u019b\7\66\2\2\u019b\u019d\7\67\2\2\u019c\u019e\7>\2\2\u019d\u019c\3"+
		"\2\2\2\u019d\u019e\3\2\2\2\u019e\u019f\3\2\2\2\u019f\u01a0\b\24\1\2\u01a0"+
		"\'\3\2\2\2\u01a1\u01a2\7$\2\2\u01a2\u01a3\7<\2\2\u01a3\u01a4\7\34\2\2"+
		"\u01a4\u01a5\7\66\2\2\u01a5\u01a6\7\35\2\2\u01a6\u01a7\7:\2\2\u01a7\u01a8"+
		"\5\66\34\2\u01a8\u01aa\7\67\2\2\u01a9\u01ab\7>\2\2\u01aa\u01a9\3\2\2\2"+
		"\u01aa\u01ab\3\2\2\2\u01ab\u01ac\3\2\2\2\u01ac\u01ad\b\25\1\2\u01ad)\3"+
		"\2\2\2\u01ae\u01af\7$\2\2\u01af\u01b0\7<\2\2\u01b0\u01b1\7\36\2\2\u01b1"+
		"\u01b2\b\26\1\2\u01b2+\3\2\2\2\u01b3\u01b4\7$\2\2\u01b4\u01b5\7<\2\2\u01b5"+
		"\u01b6\7\37\2\2\u01b6\u01b7\b\27\1\2\u01b7-\3\2\2\2\u01b8\u01b9\7$\2\2"+
		"\u01b9\u01ba\7?\2\2\u01ba\u01bb\5\66\34\2\u01bb\u01bc\7@\2\2\u01bc\u01bd"+
		"\b\30\1\2\u01bd/\3\2\2\2\u01be\u01bf\7$\2\2\u01bf\u01c0\7?\2\2\u01c0\u01c1"+
		"\5\66\34\2\u01c1\u01c2\7@\2\2\u01c2\u01c3\7*\2\2\u01c3\u01c5\5\66\34\2"+
		"\u01c4\u01c6\7>\2\2\u01c5\u01c4\3\2\2\2\u01c5\u01c6\3\2\2\2\u01c6\u01c7"+
		"\3\2\2\2\u01c7\u01c8\b\31\1\2\u01c8\61\3\2\2\2\u01c9\u01ca\7\b\2\2\u01ca"+
		"\u01cb\7$\2\2\u01cb\u01cc\7\66\2\2\u01cc\u01cd\7:\2\2\u01cd\u01ce\5<\37"+
		"\2\u01ce\u01cf\7\67\2\2\u01cf\u01d1\7*\2\2\u01d0\u01d2\7>\2\2\u01d1\u01d0"+
		"\3\2\2\2\u01d1\u01d2\3\2\2\2\u01d2\63\3\2\2\2\u01d3\u01d4\7 \2\2\u01d4"+
		"\u01d5\7$\2\2\u01d5\u01d6\7\66\2\2\u01d6\u01d7\5> \2\u01d7\u01d8\7\67"+
		"\2\2\u01d8\u01d9\7\64\2\2\u01d9\u01da\7/\2\2\u01da\u01db\5<\37\2\u01db"+
		"\u01dc\78\2\2\u01dc\u01dd\5\4\3\2\u01dd\u01de\79\2\2\u01de\u01df\b\33"+
		"\1\2\u01df\u0200\3\2\2\2\u01e0\u01e1\7 \2\2\u01e1\u01e2\7$\2\2\u01e2\u01e3"+
		"\7\66\2\2\u01e3\u01e4\5> \2\u01e4\u01e5\7\67\2\2\u01e5\u01e6\78\2\2\u01e6"+
		"\u01e7\5\4\3\2\u01e7\u01e8\79\2\2\u01e8\u01e9\b\33\1\2\u01e9\u0200\3\2"+
		"\2\2\u01ea\u01eb\7 \2\2\u01eb\u01ec\7$\2\2\u01ec\u01ed\7\66\2\2\u01ed"+
		"\u01ee\7\67\2\2\u01ee\u01ef\7\64\2\2\u01ef\u01f0\7/\2\2\u01f0\u01f1\5"+
		"<\37\2\u01f1\u01f2\78\2\2\u01f2\u01f3\5\4\3\2\u01f3\u01f4\79\2\2\u01f4"+
		"\u01f5\b\33\1\2\u01f5\u0200\3\2\2\2\u01f6\u01f7\7 \2\2\u01f7\u01f8\7$"+
		"\2\2\u01f8\u01f9\7\66\2\2\u01f9\u01fa\7\67\2\2\u01fa\u01fb\78\2\2\u01fb"+
		"\u01fc\5\4\3\2\u01fc\u01fd\79\2\2\u01fd\u01fe\b\33\1\2\u01fe\u0200\3\2"+
		"\2\2\u01ff\u01d3\3\2\2\2\u01ff\u01e0\3\2\2\2\u01ff\u01ea\3\2\2\2\u01ff"+
		"\u01f6\3\2\2\2\u0200\65\3\2\2\2\u0201\u0202\b\34\1\2\u0202\u0203\7\'\2"+
		"\2\u0203\u0204\5\66\34\16\u0204\u0205\b\34\1\2\u0205\u0226\3\2\2\2\u0206"+
		"\u0207\7\64\2\2\u0207\u0208\5\66\34\r\u0208\u0209\b\34\1\2\u0209\u0226"+
		"\3\2\2\2\u020a\u020b\7\66\2\2\u020b\u020c\5\66\34\2\u020c\u020d\7\67\2"+
		"\2\u020d\u020e\b\34\1\2\u020e\u0226\3\2\2\2\u020f\u0210\7\"\2\2\u0210"+
		"\u0226\b\34\1\2\u0211\u0212\7#\2\2\u0212\u0226\b\34\1\2\u0213\u0214\7"+
		"\n\2\2\u0214\u0226\b\34\1\2\u0215\u0216\7\13\2\2\u0216\u0226\b\34\1\2"+
		"\u0217\u0218\7\t\2\2\u0218\u0226\b\34\1\2\u0219\u021a\58\35\2\u021a\u021b"+
		"\b\34\1\2\u021b\u0226\3\2\2\2\u021c\u021d\5*\26\2\u021d\u021e\b\34\1\2"+
		"\u021e\u0226\3\2\2\2\u021f\u0220\5,\27\2\u0220\u0221\b\34\1\2\u0221\u0226"+
		"\3\2\2\2\u0222\u0223\5.\30\2\u0223\u0224\b\34\1\2\u0224\u0226\3\2\2\2"+
		"\u0225\u0201\3\2\2\2\u0225\u0206\3\2\2\2\u0225\u020a\3\2\2\2\u0225\u020f"+
		"\3\2\2\2\u0225\u0211\3\2\2\2\u0225\u0213\3\2\2\2\u0225\u0215\3\2\2\2\u0225"+
		"\u0217\3\2\2\2\u0225\u0219\3\2\2\2\u0225\u021c\3\2\2\2\u0225\u021f\3\2"+
		"\2\2\u0225\u0222\3\2\2\2\u0226\u024c\3\2\2\2\u0227\u0228\f\25\2\2\u0228"+
		"\u0229\t\3\2\2\u0229\u022a\5\66\34\26\u022a\u022b\b\34\1\2\u022b\u024b"+
		"\3\2\2\2\u022c\u022d\f\24\2\2\u022d\u022e\t\4\2\2\u022e\u022f\5\66\34"+
		"\25\u022f\u0230\b\34\1\2\u0230\u024b\3\2\2\2\u0231\u0232\f\23\2\2\u0232"+
		"\u0233\t\5\2\2\u0233\u0234\5\66\34\24\u0234\u0235\b\34\1\2\u0235\u024b"+
		"\3\2\2\2\u0236\u0237\f\22\2\2\u0237\u0238\t\6\2\2\u0238\u0239\5\66\34"+
		"\23\u0239\u023a\b\34\1\2\u023a\u024b\3\2\2\2\u023b\u023c\f\21\2\2\u023c"+
		"\u023d\t\7\2\2\u023d\u023e\5\66\34\22\u023e\u023f\b\34\1\2\u023f\u024b"+
		"\3\2\2\2\u0240\u0241\f\20\2\2\u0241\u0242\7)\2\2\u0242\u0243\5\66\34\21"+
		"\u0243\u0244\b\34\1\2\u0244\u024b\3\2\2\2\u0245\u0246\f\17\2\2\u0246\u0247"+
		"\7(\2\2\u0247\u0248\5\66\34\20\u0248\u0249\b\34\1\2\u0249\u024b\3\2\2"+
		"\2\u024a\u0227\3\2\2\2\u024a\u022c\3\2\2\2\u024a\u0231\3\2\2\2\u024a\u0236"+
		"\3\2\2\2\u024a\u023b\3\2\2\2\u024a\u0240\3\2\2\2\u024a\u0245\3\2\2\2\u024b"+
		"\u024e\3\2\2\2\u024c\u024a\3\2\2\2\u024c\u024d\3\2\2\2\u024d\67\3\2\2"+
		"\2\u024e\u024c\3\2\2\2\u024f\u0250\7$\2\2\u0250\u0251\b\35\1\2\u02519"+
		"\3\2\2\2\u0252\u0253\7$\2\2\u0253\u0254\7+\2\2\u0254\u0256\5\66\34\2\u0255"+
		"\u0257\7>\2\2\u0256\u0255\3\2\2\2\u0256\u0257\3\2\2\2\u0257\u0258\3\2"+
		"\2\2\u0258\u0259\b\36\1\2\u0259\u0263\3\2\2\2\u025a\u025b\7$\2\2\u025b"+
		"\u025c\7,\2\2\u025c\u025e\5\66\34\2\u025d\u025f\7>\2\2\u025e\u025d\3\2"+
		"\2\2\u025e\u025f\3\2\2\2\u025f\u0260\3\2\2\2\u0260\u0261\b\36\1\2\u0261"+
		"\u0263\3\2\2\2\u0262\u0252\3\2\2\2\u0262\u025a\3\2\2\2\u0263;\3\2\2\2"+
		"\u0264\u0265\7\3\2\2\u0265\u026f\b\37\1\2\u0266\u0267\7\4\2\2\u0267\u026f"+
		"\b\37\1\2\u0268\u0269\7\5\2\2\u0269\u026f\b\37\1\2\u026a\u026b\7\6\2\2"+
		"\u026b\u026f\b\37\1\2\u026c\u026d\7\7\2\2\u026d\u026f\b\37\1\2\u026e\u0264"+
		"\3\2\2\2\u026e\u0266\3\2\2\2\u026e\u0268\3\2\2\2\u026e\u026a\3\2\2\2\u026e"+
		"\u026c\3\2\2\2\u026f=\3\2\2\2\u0270\u0271\b \1\2\u0271\u0272\5\66\34\2"+
		"\u0272\u0273\b \1\2\u0273\u027b\3\2\2\2\u0274\u0275\f\4\2\2\u0275\u0276"+
		"\7=\2\2\u0276\u0277\5\66\34\2\u0277\u0278\b \1\2\u0278\u027a\3\2\2\2\u0279"+
		"\u0274\3\2\2\2\u027a\u027d\3\2\2\2\u027b\u0279\3\2\2\2\u027b\u027c\3\2"+
		"\2\2\u027c?\3\2\2\2\u027d\u027b\3\2\2\2+G~\u0085\u0090\u0099\u00a3\u00a7"+
		"\u00b0\u00b9\u00bd\u00c2\u00e1\u00f3\u0105\u010a\u0112\u0132\u013f\u0145"+
		"\u014b\u0150\u0153\u0160\u016e\u017b\u0186\u018a\u0193\u019d\u01aa\u01c5"+
		"\u01d1\u01ff\u0225\u024a\u024c\u0256\u025e\u0262\u026e\u027b";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
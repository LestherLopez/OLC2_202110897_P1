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
		RULE_removestmt = 19, RULE_emptvecstmt = 20, RULE_countvecstmt = 21, RULE_accessvecstmt = 22, 
		RULE_assignationvecstmt = 23, RULE_declarematrixstmt = 24, RULE_listParams = 25, 
		RULE_expr = 26, RULE_accessstmt = 27, RULE_increaseanddecreasestmt = 28, 
		RULE_type = 29;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "block", "instruction", "printstmt", "declarestmt", "constantstmt", 
			"blockelifs", "ifstmt", "switchstmt", "casestmt", "blockcases", "assignationstmt", 
			"whilestmt", "forstmt", "guardstmt", "transferstmt", "declarevectorstmt", 
			"appendstmt", "removelaststmt", "removestmt", "emptvecstmt", "countvecstmt", 
			"accessvecstmt", "assignationvecstmt", "declarematrixstmt", "listParams", 
			"expr", "accessstmt", "increaseanddecreasestmt", "type"
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
			setState(60);
			((SContext)_localctx).block = block();
			setState(61);
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
			setState(65); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(64);
					((BlockContext)_localctx).instruction = instruction();
					((BlockContext)_localctx).ins.add(((BlockContext)_localctx).instruction);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(67); 
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
		public InstructionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruction; }
	}

	public final InstructionContext instruction() throws RecognitionException {
		InstructionContext _localctx = new InstructionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_instruction);
		try {
			setState(119);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(71);
				((InstructionContext)_localctx).printstmt = printstmt();
				 _localctx.inst = ((InstructionContext)_localctx).printstmt.prnt
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(74);
				((InstructionContext)_localctx).declarestmt = declarestmt();
				_localctx.inst = ((InstructionContext)_localctx).declarestmt.dec
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(77);
				((InstructionContext)_localctx).constantstmt = constantstmt();
				_localctx.inst = ((InstructionContext)_localctx).constantstmt.const
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(80);
				((InstructionContext)_localctx).assignationstmt = assignationstmt();
				_localctx.inst = ((InstructionContext)_localctx).assignationstmt.assign
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(83);
				((InstructionContext)_localctx).increaseanddecreasestmt = increaseanddecreasestmt();
				_localctx.inst = ((InstructionContext)_localctx).increaseanddecreasestmt.increasedecrease
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(86);
				((InstructionContext)_localctx).ifstmt = ifstmt();
				 _localctx.inst = ((InstructionContext)_localctx).ifstmt.ift 
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(89);
				((InstructionContext)_localctx).whilestmt = whilestmt();
				_localctx.inst = ((InstructionContext)_localctx).whilestmt.while
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(92);
				((InstructionContext)_localctx).forstmt = forstmt();
				_localctx.inst = ((InstructionContext)_localctx).forstmt.for
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(95);
				((InstructionContext)_localctx).switchstmt = switchstmt();
				_localctx.inst = ((InstructionContext)_localctx).switchstmt.switch
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(98);
				((InstructionContext)_localctx).guardstmt = guardstmt();
				_localctx.inst = ((InstructionContext)_localctx).guardstmt.gua
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(101);
				((InstructionContext)_localctx).transferstmt = transferstmt();
				_localctx.inst = ((InstructionContext)_localctx).transferstmt.tran
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(104);
				((InstructionContext)_localctx).declarevectorstmt = declarevectorstmt();
				_localctx.inst = ((InstructionContext)_localctx).declarevectorstmt.decvec
				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(107);
				((InstructionContext)_localctx).appendstmt = appendstmt();
				_localctx.inst = ((InstructionContext)_localctx).appendstmt.app
				}
				break;
			case 14:
				enterOuterAlt(_localctx, 14);
				{
				setState(110);
				((InstructionContext)_localctx).removelaststmt = removelaststmt();
				_localctx.inst = ((InstructionContext)_localctx).removelaststmt.removl
				}
				break;
			case 15:
				enterOuterAlt(_localctx, 15);
				{
				setState(113);
				((InstructionContext)_localctx).removestmt = removestmt();
				_localctx.inst = ((InstructionContext)_localctx).removestmt.remov
				}
				break;
			case 16:
				enterOuterAlt(_localctx, 16);
				{
				setState(116);
				((InstructionContext)_localctx).assignationvecstmt = assignationvecstmt();
				_localctx.inst = ((InstructionContext)_localctx).assignationvecstmt.assignvec
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
			setState(121);
			((PrintstmtContext)_localctx).PRINT = match(PRINT);
			setState(122);
			match(PARIZQ);
			setState(123);
			((PrintstmtContext)_localctx).expr = expr(0);
			setState(124);
			match(PARDER);
			setState(126);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PTCOMA) {
				{
				setState(125);
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
			setState(160);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(130);
				((DeclarestmtContext)_localctx).VAR = match(VAR);
				setState(131);
				((DeclarestmtContext)_localctx).ID = match(ID);
				setState(132);
				match(DOUBLEPTS);
				setState(133);
				((DeclarestmtContext)_localctx).type = type();
				setState(134);
				match(IG);
				setState(135);
				((DeclarestmtContext)_localctx).expr = expr(0);
				setState(137);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(136);
					match(PTCOMA);
					}
				}

				_localctx.dec = instructions.NewTodeclare((((DeclarestmtContext)_localctx).VAR!=null?((DeclarestmtContext)_localctx).VAR.getLine():0), (((DeclarestmtContext)_localctx).VAR!=null?((DeclarestmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarestmtContext)_localctx).ID!=null?((DeclarestmtContext)_localctx).ID.getText():null), ((DeclarestmtContext)_localctx).type.t, ((DeclarestmtContext)_localctx).expr.e, false)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(141);
				((DeclarestmtContext)_localctx).VAR = match(VAR);
				setState(142);
				((DeclarestmtContext)_localctx).ID = match(ID);
				setState(143);
				match(IG);
				setState(144);
				((DeclarestmtContext)_localctx).expr = expr(0);
				setState(146);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(145);
					match(PTCOMA);
					}
				}

				_localctx.dec = instructions.NewTodeclare((((DeclarestmtContext)_localctx).VAR!=null?((DeclarestmtContext)_localctx).VAR.getLine():0), (((DeclarestmtContext)_localctx).VAR!=null?((DeclarestmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarestmtContext)_localctx).ID!=null?((DeclarestmtContext)_localctx).ID.getText():null), environment.NULL, ((DeclarestmtContext)_localctx).expr.e, false)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(150);
				((DeclarestmtContext)_localctx).VAR = match(VAR);
				setState(151);
				((DeclarestmtContext)_localctx).ID = match(ID);
				setState(152);
				match(DOUBLEPTS);
				setState(153);
				((DeclarestmtContext)_localctx).type = type();
				setState(154);
				match(QUESTION);
				setState(156);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(155);
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
			setState(182);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(162);
				((ConstantstmtContext)_localctx).LET = match(LET);
				setState(163);
				((ConstantstmtContext)_localctx).ID = match(ID);
				setState(164);
				match(DOUBLEPTS);
				setState(165);
				((ConstantstmtContext)_localctx).type = type();
				setState(166);
				match(IG);
				setState(167);
				((ConstantstmtContext)_localctx).expr = expr(0);
				setState(169);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(168);
					match(PTCOMA);
					}
				}

				_localctx.const = instructions.NewTodeclare((((ConstantstmtContext)_localctx).LET!=null?((ConstantstmtContext)_localctx).LET.getLine():0), (((ConstantstmtContext)_localctx).LET!=null?((ConstantstmtContext)_localctx).LET.getCharPositionInLine():0), (((ConstantstmtContext)_localctx).ID!=null?((ConstantstmtContext)_localctx).ID.getText():null), ((ConstantstmtContext)_localctx).type.t, ((ConstantstmtContext)_localctx).expr.e, true)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(173);
				((ConstantstmtContext)_localctx).LET = match(LET);
				setState(174);
				((ConstantstmtContext)_localctx).ID = match(ID);
				setState(175);
				match(IG);
				setState(176);
				((ConstantstmtContext)_localctx).expr = expr(0);
				setState(178);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(177);
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
			setState(185); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(184);
					((BlockelifsContext)_localctx).ifstmt = ifstmt();
					((BlockelifsContext)_localctx).elif.add(((BlockelifsContext)_localctx).ifstmt);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(187); 
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
			setState(218);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(191);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(192);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(193);
				match(LLAVEIZQ);
				setState(194);
				((IfstmtContext)_localctx).ifb = block();
				setState(195);
				match(LLAVEDER);
				 ((IfstmtContext)_localctx).ift =  instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).ifb.blk, nil); 
				                                        _localctx.el = ((IfstmtContext)_localctx).ifb.blk
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(198);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(199);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(200);
				match(LLAVEIZQ);
				setState(201);
				((IfstmtContext)_localctx).ifelseblck = block();
				setState(202);
				match(LLAVEDER);
				setState(203);
				match(ELSE);
				setState(204);
				match(LLAVEIZQ);
				setState(205);
				((IfstmtContext)_localctx).elseifblck = block();
				setState(206);
				match(LLAVEDER);
				 ((IfstmtContext)_localctx).ift =  instructions.NewIf((((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getLine():0), (((IfstmtContext)_localctx).IF!=null?((IfstmtContext)_localctx).IF.getCharPositionInLine():0), ((IfstmtContext)_localctx).expr.e, ((IfstmtContext)_localctx).ifelseblck.blk, ((IfstmtContext)_localctx).elseifblck.blk); 
				                                                                                        ((IfstmtContext)_localctx).el =  ((IfstmtContext)_localctx).ifelseblck.blk;
				                                                                                        
				                                                                                        
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(209);
				((IfstmtContext)_localctx).IF = match(IF);
				setState(210);
				((IfstmtContext)_localctx).expr = expr(0);
				setState(211);
				match(LLAVEIZQ);
				setState(212);
				((IfstmtContext)_localctx).elif = block();
				setState(213);
				match(LLAVEDER);
				setState(214);
				match(ELSE);
				setState(215);
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
			setState(236);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(220);
				((SwitchstmtContext)_localctx).SWITCH = match(SWITCH);
				setState(221);
				((SwitchstmtContext)_localctx).expr = expr(0);
				setState(222);
				match(LLAVEIZQ);
				setState(223);
				match(DEFAULT);
				setState(224);
				match(DOUBLEPTS);
				setState(225);
				((SwitchstmtContext)_localctx).block = block();
				setState(226);
				match(LLAVEDER);
				 ((SwitchstmtContext)_localctx).switch =  instructions.NewSwitch((((SwitchstmtContext)_localctx).SWITCH!=null?((SwitchstmtContext)_localctx).SWITCH.getLine():0), (((SwitchstmtContext)_localctx).SWITCH!=null?((SwitchstmtContext)_localctx).SWITCH.getCharPositionInLine():0), ((SwitchstmtContext)_localctx).expr.e, nil,  nil, ((SwitchstmtContext)_localctx).block.blk);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(229);
				((SwitchstmtContext)_localctx).SWITCH = match(SWITCH);
				setState(230);
				((SwitchstmtContext)_localctx).expr = expr(0);
				setState(231);
				match(LLAVEIZQ);
				setState(232);
				((SwitchstmtContext)_localctx).casestmt = casestmt();
				setState(233);
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
			setState(254);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(238);
				((CasestmtContext)_localctx).CASE = match(CASE);
				setState(239);
				((CasestmtContext)_localctx).expr = expr(0);
				setState(240);
				match(DOUBLEPTS);
				setState(241);
				((CasestmtContext)_localctx).block = block();
				setState(242);
				((CasestmtContext)_localctx).blockcases = blockcases();
				 _localctx.cas = instructions.NewCase((((CasestmtContext)_localctx).CASE!=null?((CasestmtContext)_localctx).CASE.getLine():0),(((CasestmtContext)_localctx).CASE!=null?((CasestmtContext)_localctx).CASE.getCharPositionInLine():0), ((CasestmtContext)_localctx).expr.e,((CasestmtContext)_localctx).block.blk, ((CasestmtContext)_localctx).blockcases.blkcase)
				                                fmt.Println("entro")
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(245);
				((CasestmtContext)_localctx).CASE = match(CASE);
				setState(246);
				((CasestmtContext)_localctx).expr = expr(0);
				setState(247);
				match(DOUBLEPTS);
				setState(248);
				((CasestmtContext)_localctx).sen = block();
				setState(249);
				match(DEFAULT);
				setState(250);
				match(DOUBLEPTS);
				setState(251);
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
			setState(257); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(256);
					((BlockcasesContext)_localctx).casestmt = casestmt();
					((BlockcasesContext)_localctx).casedef.add(((BlockcasesContext)_localctx).casestmt);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(259); 
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
			setState(263);
			((AssignationstmtContext)_localctx).ID = match(ID);
			setState(264);
			((AssignationstmtContext)_localctx).IG = match(IG);
			setState(265);
			((AssignationstmtContext)_localctx).expr = expr(0);
			setState(267);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PTCOMA) {
				{
				setState(266);
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
			setState(271);
			((WhilestmtContext)_localctx).WHILE = match(WHILE);
			setState(272);
			((WhilestmtContext)_localctx).expr = expr(0);
			setState(273);
			match(LLAVEIZQ);
			setState(274);
			((WhilestmtContext)_localctx).block = block();
			setState(275);
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
			setState(299);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(278);
				((ForstmtContext)_localctx).FOR = match(FOR);
				setState(279);
				((ForstmtContext)_localctx).ID = match(ID);
				setState(280);
				match(IN);
				setState(281);
				((ForstmtContext)_localctx).expr = expr(0);
				setState(282);
				match(LLAVEIZQ);
				setState(283);
				((ForstmtContext)_localctx).block = block();
				setState(284);
				match(LLAVEDER);
				_localctx.for = instructions.NewFor((((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getLine():0), (((ForstmtContext)_localctx).FOR!=null?((ForstmtContext)_localctx).FOR.getCharPositionInLine():0), (((ForstmtContext)_localctx).ID!=null?((ForstmtContext)_localctx).ID.getText():null), ((ForstmtContext)_localctx).expr.e, nil, ((ForstmtContext)_localctx).block.blk)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(287);
				((ForstmtContext)_localctx).FOR = match(FOR);
				setState(288);
				((ForstmtContext)_localctx).ID = match(ID);
				setState(289);
				match(IN);
				setState(290);
				((ForstmtContext)_localctx).first = expr(0);
				setState(291);
				match(POINT);
				setState(292);
				match(POINT);
				setState(293);
				((ForstmtContext)_localctx).second = expr(0);
				setState(294);
				match(LLAVEIZQ);
				setState(295);
				((ForstmtContext)_localctx).block = block();
				setState(296);
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
			setState(301);
			((GuardstmtContext)_localctx).GUARD = match(GUARD);
			setState(302);
			((GuardstmtContext)_localctx).expr = expr(0);
			setState(303);
			match(ELSE);
			setState(304);
			match(LLAVEIZQ);
			setState(305);
			((GuardstmtContext)_localctx).block = block();
			setState(306);
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
			setState(307);
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
			setState(332);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,21,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(310);
				((TransferstmtContext)_localctx).RETURN = match(RETURN);
				setState(312);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(311);
					match(PTCOMA);
					}
				}

				_localctx.tran = instructions.NewReturnIn((((TransferstmtContext)_localctx).RETURN!=null?((TransferstmtContext)_localctx).RETURN.getLine():0), (((TransferstmtContext)_localctx).RETURN!=null?((TransferstmtContext)_localctx).RETURN.getCharPositionInLine():0), nil)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(315);
				((TransferstmtContext)_localctx).RETURN = match(RETURN);
				setState(316);
				((TransferstmtContext)_localctx).expr = expr(0);
				setState(318);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(317);
					match(PTCOMA);
					}
				}

				_localctx.tran = instructions.NewReturnIn((((TransferstmtContext)_localctx).RETURN!=null?((TransferstmtContext)_localctx).RETURN.getLine():0), (((TransferstmtContext)_localctx).RETURN!=null?((TransferstmtContext)_localctx).RETURN.getCharPositionInLine():0), ((TransferstmtContext)_localctx).expr.e)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(322);
				((TransferstmtContext)_localctx).CONTINUE = match(CONTINUE);
				setState(324);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(323);
					match(PTCOMA);
					}
				}

				_localctx.tran = instructions.NewContinue((((TransferstmtContext)_localctx).CONTINUE!=null?((TransferstmtContext)_localctx).CONTINUE.getLine():0), (((TransferstmtContext)_localctx).CONTINUE!=null?((TransferstmtContext)_localctx).CONTINUE.getCharPositionInLine():0))
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(327);
				((TransferstmtContext)_localctx).BREAK = match(BREAK);
				setState(329);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(328);
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
			setState(387);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,26,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(334);
				((DeclarevectorstmtContext)_localctx).VAR = match(VAR);
				setState(335);
				((DeclarevectorstmtContext)_localctx).ID = match(ID);
				setState(336);
				match(DOUBLEPTS);
				setState(337);
				match(CORCHETEIZQ);
				setState(338);
				((DeclarevectorstmtContext)_localctx).type = type();
				setState(339);
				match(CORCHETEDER);
				setState(340);
				match(IG);
				setState(341);
				match(CORCHETEIZQ);
				setState(342);
				((DeclarevectorstmtContext)_localctx).listParams = listParams(0);
				setState(343);
				match(CORCHETEDER);
				setState(345);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(344);
					match(PTCOMA);
					}
				}

				_localctx.decvec = instructions.NewToDecalreVector((((DeclarevectorstmtContext)_localctx).VAR!=null?((DeclarevectorstmtContext)_localctx).VAR.getLine():0), (((DeclarevectorstmtContext)_localctx).VAR!=null?((DeclarevectorstmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarevectorstmtContext)_localctx).ID!=null?((DeclarevectorstmtContext)_localctx).ID.getText():null), ((DeclarevectorstmtContext)_localctx).type.t, ((DeclarevectorstmtContext)_localctx).listParams.l, "")
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(349);
				((DeclarevectorstmtContext)_localctx).VAR = match(VAR);
				setState(350);
				((DeclarevectorstmtContext)_localctx).ID = match(ID);
				setState(351);
				match(DOUBLEPTS);
				setState(352);
				match(CORCHETEIZQ);
				setState(353);
				((DeclarevectorstmtContext)_localctx).type = type();
				setState(354);
				match(CORCHETEDER);
				setState(355);
				match(IG);
				setState(356);
				match(CORCHETEIZQ);
				setState(357);
				match(CORCHETEDER);
				setState(359);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(358);
					match(PTCOMA);
					}
				}

				_localctx.decvec = instructions.NewToDecalreVector((((DeclarevectorstmtContext)_localctx).VAR!=null?((DeclarevectorstmtContext)_localctx).VAR.getLine():0), (((DeclarevectorstmtContext)_localctx).VAR!=null?((DeclarevectorstmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarevectorstmtContext)_localctx).ID!=null?((DeclarevectorstmtContext)_localctx).ID.getText():null), ((DeclarevectorstmtContext)_localctx).type.t, nil, "")
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(363);
				match(VAR);
				setState(364);
				match(ID);
				setState(365);
				match(IG);
				setState(366);
				match(CORCHETEIZQ);
				setState(367);
				type();
				setState(368);
				match(CORCHETEDER);
				setState(369);
				match(PARIZQ);
				setState(370);
				match(PARDER);
				setState(372);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(371);
					match(PTCOMA);
					}
				}

				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(374);
				((DeclarevectorstmtContext)_localctx).VAR = match(VAR);
				setState(375);
				((DeclarevectorstmtContext)_localctx).F = match(ID);
				setState(376);
				match(DOUBLEPTS);
				setState(377);
				match(CORCHETEIZQ);
				setState(378);
				((DeclarevectorstmtContext)_localctx).type = type();
				setState(379);
				match(CORCHETEDER);
				setState(380);
				match(IG);
				setState(381);
				((DeclarevectorstmtContext)_localctx).S = match(ID);
				setState(383);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(382);
					match(PTCOMA);
					}
				}

				_localctx.decvec = instructions.NewToDecalreVector((((DeclarevectorstmtContext)_localctx).VAR!=null?((DeclarevectorstmtContext)_localctx).VAR.getLine():0), (((DeclarevectorstmtContext)_localctx).VAR!=null?((DeclarevectorstmtContext)_localctx).VAR.getCharPositionInLine():0), (((DeclarevectorstmtContext)_localctx).F!=null?((DeclarevectorstmtContext)_localctx).F.getText():null), ((DeclarevectorstmtContext)_localctx).type.t, nil, (((DeclarevectorstmtContext)_localctx).S!=null?((DeclarevectorstmtContext)_localctx).S.getText():null))
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
			setState(389);
			((AppendstmtContext)_localctx).ID = match(ID);
			setState(390);
			match(POINT);
			setState(391);
			match(APPEND);
			setState(392);
			match(PARIZQ);
			setState(393);
			((AppendstmtContext)_localctx).expr = expr(0);
			setState(394);
			match(PARDER);
			setState(396);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PTCOMA) {
				{
				setState(395);
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
			setState(400);
			((RemovelaststmtContext)_localctx).ID = match(ID);
			setState(401);
			match(POINT);
			setState(402);
			match(REMOVELAST);
			setState(403);
			match(PARIZQ);
			setState(404);
			match(PARDER);
			setState(406);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PTCOMA) {
				{
				setState(405);
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
			setState(410);
			((RemovestmtContext)_localctx).ID = match(ID);
			setState(411);
			match(POINT);
			setState(412);
			match(REMOVE);
			setState(413);
			match(PARIZQ);
			setState(414);
			match(AT);
			setState(415);
			match(DOUBLEPTS);
			setState(416);
			((RemovestmtContext)_localctx).expr = expr(0);
			setState(417);
			match(PARDER);
			setState(419);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PTCOMA) {
				{
				setState(418);
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
			setState(423);
			((EmptvecstmtContext)_localctx).ID = match(ID);
			setState(424);
			match(POINT);
			setState(425);
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
			setState(428);
			((CountvecstmtContext)_localctx).ID = match(ID);
			setState(429);
			match(POINT);
			setState(430);
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
			setState(433);
			((AccessvecstmtContext)_localctx).ID = match(ID);
			setState(434);
			match(CORCHETEIZQ);
			setState(435);
			((AccessvecstmtContext)_localctx).expr = expr(0);
			setState(436);
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
			setState(439);
			((AssignationvecstmtContext)_localctx).ID = match(ID);
			setState(440);
			match(CORCHETEIZQ);
			setState(441);
			((AssignationvecstmtContext)_localctx).expprim = expr(0);
			setState(442);
			match(CORCHETEDER);
			setState(443);
			match(IG);
			setState(444);
			((AssignationvecstmtContext)_localctx).expsegundo = expr(0);
			setState(446);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PTCOMA) {
				{
				setState(445);
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
			setState(450);
			match(VAR);
			setState(451);
			match(ID);
			setState(452);
			match(PARIZQ);
			setState(453);
			match(DOUBLEPTS);
			setState(454);
			type();
			setState(455);
			match(PARDER);
			setState(456);
			match(IG);
			setState(458);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==PTCOMA) {
				{
				setState(457);
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
		int _startState = 50;
		enterRecursionRule(_localctx, 50, RULE_listParams, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(461);
			((ListParamsContext)_localctx).expr = expr(0);

			            _localctx.l = []interface{}{}
			            _localctx.l = append(_localctx.l, ((ListParamsContext)_localctx).expr.e)
			        
			}
			_ctx.stop = _input.LT(-1);
			setState(471);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,32,_ctx);
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
					setState(464);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(465);
					match(COMA);
					setState(466);
					((ListParamsContext)_localctx).expr = expr(0);

					                                          var arr []interface{}
					                                          arr = append(((ListParamsContext)_localctx).list.l, ((ListParamsContext)_localctx).expr.e)
					                                          _localctx.l = arr
					                                      
					}
					} 
				}
				setState(473);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,32,_ctx);
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
			setState(510);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,33,_ctx) ) {
			case 1:
				{
				setState(475);
				((ExprContext)_localctx).op = match(NOT);
				setState(476);
				((ExprContext)_localctx).left = ((ExprContext)_localctx).expr = expr(12);
				 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).left.e) 
				}
				break;
			case 2:
				{
				setState(479);
				((ExprContext)_localctx).op = match(SUB);
				setState(480);
				((ExprContext)_localctx).left = ((ExprContext)_localctx).expr = expr(11);
				 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e,  (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), nil)
				}
				break;
			case 3:
				{
				setState(483);
				match(PARIZQ);
				setState(484);
				((ExprContext)_localctx).expr = expr(0);
				setState(485);
				match(PARDER);
				 _localctx.e = ((ExprContext)_localctx).expr.e 
				}
				break;
			case 4:
				{
				setState(488);
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
				setState(490);
				((ExprContext)_localctx).STRING = match(STRING);

				        str := (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getText():null)
				        _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getLine():0), (((ExprContext)_localctx).STRING!=null?((ExprContext)_localctx).STRING.getCharPositionInLine():0), str[1:len(str)-1],environment.STRING)
				    
				}
				break;
			case 6:
				{
				setState(492);
				((ExprContext)_localctx).TRUE = match(TRUE);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).TRUE!=null?((ExprContext)_localctx).TRUE.getLine():0), (((ExprContext)_localctx).TRUE!=null?((ExprContext)_localctx).TRUE.getCharPositionInLine():0), true, environment.BOOLEAN) 
				}
				break;
			case 7:
				{
				setState(494);
				((ExprContext)_localctx).FALSE = match(FALSE);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).FALSE!=null?((ExprContext)_localctx).FALSE.getLine():0), (((ExprContext)_localctx).FALSE!=null?((ExprContext)_localctx).FALSE.getCharPositionInLine():0), false, environment.BOOLEAN) 
				}
				break;
			case 8:
				{
				setState(496);
				((ExprContext)_localctx).NIL = match(NIL);
				 _localctx.e = expressions.NewPrimitive((((ExprContext)_localctx).NIL!=null?((ExprContext)_localctx).NIL.getLine():0), (((ExprContext)_localctx).NIL!=null?((ExprContext)_localctx).NIL.getCharPositionInLine():0), nil, environment.NULL) 
				}
				break;
			case 9:
				{
				setState(498);
				((ExprContext)_localctx).accessstmt = accessstmt();
				_localctx.e = ((ExprContext)_localctx).accessstmt.access
				}
				break;
			case 10:
				{
				setState(501);
				((ExprContext)_localctx).emptvecstmt = emptvecstmt();
				_localctx.e = ((ExprContext)_localctx).emptvecstmt.emptyvec
				}
				break;
			case 11:
				{
				setState(504);
				((ExprContext)_localctx).countvecstmt = countvecstmt();
				_localctx.e = ((ExprContext)_localctx).countvecstmt.count
				}
				break;
			case 12:
				{
				setState(507);
				((ExprContext)_localctx).accessvecstmt = accessvecstmt();
				_localctx.e = ((ExprContext)_localctx).accessvecstmt.accessvec
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(549);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,35,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(547);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,34,_ctx) ) {
					case 1:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.left = _prevctx;
						_localctx.left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(512);
						if (!(precpred(_ctx, 19))) throw new FailedPredicateException(this, "precpred(_ctx, 19)");
						setState(513);
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
						setState(514);
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
						setState(517);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(518);
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
						setState(519);
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
						setState(522);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(523);
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
						setState(524);
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
						setState(527);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(528);
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
						setState(529);
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
						setState(532);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(533);
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
						setState(534);
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
						setState(537);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(538);
						((ExprContext)_localctx).op = match(AND);
						setState(539);
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
						setState(542);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(543);
						((ExprContext)_localctx).op = match(OR);
						setState(544);
						((ExprContext)_localctx).right = ((ExprContext)_localctx).expr = expr(14);
						 _localctx.e = expressions.NewOperation((((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetLine(), (((ExprContext)_localctx).left!=null?(((ExprContext)_localctx).left.start):null).GetColumn(), ((ExprContext)_localctx).left.e, (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null), ((ExprContext)_localctx).right.e) 
						}
						break;
					}
					} 
				}
				setState(551);
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
			setState(552);
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
			setState(571);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,38,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(555);
				((IncreaseanddecreasestmtContext)_localctx).ID = match(ID);
				setState(556);
				((IncreaseanddecreasestmtContext)_localctx).IG_ADD = match(IG_ADD);
				setState(557);
				((IncreaseanddecreasestmtContext)_localctx).expr = expr(0);
				setState(559);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(558);
					match(PTCOMA);
					}
				}

				_localctx.increasedecrease = instructions.NewIncreaseDecrease((((IncreaseanddecreasestmtContext)_localctx).ID!=null?((IncreaseanddecreasestmtContext)_localctx).ID.getLine():0), (((IncreaseanddecreasestmtContext)_localctx).ID!=null?((IncreaseanddecreasestmtContext)_localctx).ID.getCharPositionInLine():0), (((IncreaseanddecreasestmtContext)_localctx).ID!=null?((IncreaseanddecreasestmtContext)_localctx).ID.getText():null), (((IncreaseanddecreasestmtContext)_localctx).IG_ADD!=null?((IncreaseanddecreasestmtContext)_localctx).IG_ADD.getText():null), ((IncreaseanddecreasestmtContext)_localctx).expr.e)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(563);
				((IncreaseanddecreasestmtContext)_localctx).ID = match(ID);
				setState(564);
				((IncreaseanddecreasestmtContext)_localctx).IG_SUB = match(IG_SUB);
				setState(565);
				((IncreaseanddecreasestmtContext)_localctx).expr = expr(0);
				setState(567);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(566);
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
			setState(583);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRINGS:
				enterOuterAlt(_localctx, 1);
				{
				setState(573);
				match(STRINGS);
				_localctx.t = environment.STRING
				}
				break;
			case INTS:
				enterOuterAlt(_localctx, 2);
				{
				setState(575);
				match(INTS);
				_localctx.t = environment.INTEGER
				}
				break;
			case FLOATS:
				enterOuterAlt(_localctx, 3);
				{
				setState(577);
				match(FLOATS);
				_localctx.t = environment.FLOAT
				}
				break;
			case BOOLS:
				enterOuterAlt(_localctx, 4);
				{
				setState(579);
				match(BOOLS);
				_localctx.t = environment.BOOLEAN 
				}
				break;
			case CHARACTERS:
				enterOuterAlt(_localctx, 5);
				{
				setState(581);
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
		case 25:
			return listParams_sempred((ListParamsContext)_localctx, predIndex);
		case 26:
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
			return precpred(_ctx, 19);
		case 2:
			return precpred(_ctx, 18);
		case 3:
			return precpred(_ctx, 17);
		case 4:
			return precpred(_ctx, 16);
		case 5:
			return precpred(_ctx, 15);
		case 6:
			return precpred(_ctx, 14);
		case 7:
			return precpred(_ctx, 13);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3A\u024c\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\3\2\3\2\3"+
		"\2\3\2\3\3\6\3D\n\3\r\3\16\3E\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\5\4z\n\4\3\5\3\5\3\5\3\5\3\5\5\5\u0081\n\5\3\5\3"+
		"\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6\u008c\n\6\3\6\3\6\3\6\3\6\3\6\3\6\3"+
		"\6\5\6\u0095\n\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6\u009f\n\6\3\6\3\6"+
		"\5\6\u00a3\n\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\5\7\u00ac\n\7\3\7\3\7\3\7\3"+
		"\7\3\7\3\7\3\7\5\7\u00b5\n\7\3\7\3\7\5\7\u00b9\n\7\3\b\6\b\u00bc\n\b\r"+
		"\b\16\b\u00bd\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t"+
		"\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\5\t\u00dd"+
		"\n\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\5"+
		"\n\u00ef\n\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3"+
		"\13\3\13\3\13\3\13\3\13\5\13\u0101\n\13\3\f\6\f\u0104\n\f\r\f\16\f\u0105"+
		"\3\f\3\f\3\r\3\r\3\r\3\r\5\r\u010e\n\r\3\r\3\r\3\16\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17"+
		"\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\5\17\u012e\n\17\3\20\3\20"+
		"\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\21\3\21\5\21\u013b\n\21\3\21\3\21"+
		"\3\21\3\21\5\21\u0141\n\21\3\21\3\21\3\21\3\21\5\21\u0147\n\21\3\21\3"+
		"\21\3\21\5\21\u014c\n\21\3\21\5\21\u014f\n\21\3\22\3\22\3\22\3\22\3\22"+
		"\3\22\3\22\3\22\3\22\3\22\3\22\5\22\u015c\n\22\3\22\3\22\3\22\3\22\3\22"+
		"\3\22\3\22\3\22\3\22\3\22\3\22\3\22\5\22\u016a\n\22\3\22\3\22\3\22\3\22"+
		"\3\22\3\22\3\22\3\22\3\22\3\22\3\22\5\22\u0177\n\22\3\22\3\22\3\22\3\22"+
		"\3\22\3\22\3\22\3\22\3\22\5\22\u0182\n\22\3\22\3\22\5\22\u0186\n\22\3"+
		"\23\3\23\3\23\3\23\3\23\3\23\3\23\5\23\u018f\n\23\3\23\3\23\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\5\24\u0199\n\24\3\24\3\24\3\25\3\25\3\25\3\25\3\25"+
		"\3\25\3\25\3\25\3\25\5\25\u01a6\n\25\3\25\3\25\3\26\3\26\3\26\3\26\3\26"+
		"\3\27\3\27\3\27\3\27\3\27\3\30\3\30\3\30\3\30\3\30\3\30\3\31\3\31\3\31"+
		"\3\31\3\31\3\31\3\31\5\31\u01c1\n\31\3\31\3\31\3\32\3\32\3\32\3\32\3\32"+
		"\3\32\3\32\3\32\5\32\u01cd\n\32\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\33"+
		"\3\33\7\33\u01d8\n\33\f\33\16\33\u01db\13\33\3\34\3\34\3\34\3\34\3\34"+
		"\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34"+
		"\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34"+
		"\3\34\3\34\3\34\5\34\u0201\n\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34"+
		"\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34"+
		"\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\7\34"+
		"\u0226\n\34\f\34\16\34\u0229\13\34\3\35\3\35\3\35\3\36\3\36\3\36\3\36"+
		"\5\36\u0232\n\36\3\36\3\36\3\36\3\36\3\36\3\36\5\36\u023a\n\36\3\36\3"+
		"\36\5\36\u023e\n\36\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37"+
		"\5\37\u024a\n\37\3\37\2\4\64\66 \2\4\6\b\n\f\16\20\22\24\26\30\32\34\36"+
		" \"$&(*,.\60\62\64\668:<\2\b\3\2\27\31\3\2/\60\3\2\61\63\4\2++--\4\2,"+
		",..\3\2#$\2\u027b\2>\3\2\2\2\4C\3\2\2\2\6y\3\2\2\2\b{\3\2\2\2\n\u00a2"+
		"\3\2\2\2\f\u00b8\3\2\2\2\16\u00bb\3\2\2\2\20\u00dc\3\2\2\2\22\u00ee\3"+
		"\2\2\2\24\u0100\3\2\2\2\26\u0103\3\2\2\2\30\u0109\3\2\2\2\32\u0111\3\2"+
		"\2\2\34\u012d\3\2\2\2\36\u012f\3\2\2\2 \u014e\3\2\2\2\"\u0185\3\2\2\2"+
		"$\u0187\3\2\2\2&\u0192\3\2\2\2(\u019c\3\2\2\2*\u01a9\3\2\2\2,\u01ae\3"+
		"\2\2\2.\u01b3\3\2\2\2\60\u01b9\3\2\2\2\62\u01c4\3\2\2\2\64\u01ce\3\2\2"+
		"\2\66\u0200\3\2\2\28\u022a\3\2\2\2:\u023d\3\2\2\2<\u0249\3\2\2\2>?\5\4"+
		"\3\2?@\7\2\2\3@A\b\2\1\2A\3\3\2\2\2BD\5\6\4\2CB\3\2\2\2DE\3\2\2\2EC\3"+
		"\2\2\2EF\3\2\2\2FG\3\2\2\2GH\b\3\1\2H\5\3\2\2\2IJ\5\b\5\2JK\b\4\1\2Kz"+
		"\3\2\2\2LM\5\n\6\2MN\b\4\1\2Nz\3\2\2\2OP\5\f\7\2PQ\b\4\1\2Qz\3\2\2\2R"+
		"S\5\30\r\2ST\b\4\1\2Tz\3\2\2\2UV\5:\36\2VW\b\4\1\2Wz\3\2\2\2XY\5\20\t"+
		"\2YZ\b\4\1\2Zz\3\2\2\2[\\\5\32\16\2\\]\b\4\1\2]z\3\2\2\2^_\5\34\17\2_"+
		"`\b\4\1\2`z\3\2\2\2ab\5\22\n\2bc\b\4\1\2cz\3\2\2\2de\5\36\20\2ef\b\4\1"+
		"\2fz\3\2\2\2gh\5 \21\2hi\b\4\1\2iz\3\2\2\2jk\5\"\22\2kl\b\4\1\2lz\3\2"+
		"\2\2mn\5$\23\2no\b\4\1\2oz\3\2\2\2pq\5&\24\2qr\b\4\1\2rz\3\2\2\2st\5("+
		"\25\2tu\b\4\1\2uz\3\2\2\2vw\5\60\31\2wx\b\4\1\2xz\3\2\2\2yI\3\2\2\2yL"+
		"\3\2\2\2yO\3\2\2\2yR\3\2\2\2yU\3\2\2\2yX\3\2\2\2y[\3\2\2\2y^\3\2\2\2y"+
		"a\3\2\2\2yd\3\2\2\2yg\3\2\2\2yj\3\2\2\2ym\3\2\2\2yp\3\2\2\2ys\3\2\2\2"+
		"yv\3\2\2\2z\7\3\2\2\2{|\7\f\2\2|}\7\64\2\2}~\5\66\34\2~\u0080\7\65\2\2"+
		"\177\u0081\7<\2\2\u0080\177\3\2\2\2\u0080\u0081\3\2\2\2\u0081\u0082\3"+
		"\2\2\2\u0082\u0083\b\5\1\2\u0083\t\3\2\2\2\u0084\u0085\7\b\2\2\u0085\u0086"+
		"\7\"\2\2\u0086\u0087\78\2\2\u0087\u0088\5<\37\2\u0088\u0089\7(\2\2\u0089"+
		"\u008b\5\66\34\2\u008a\u008c\7<\2\2\u008b\u008a\3\2\2\2\u008b\u008c\3"+
		"\2\2\2\u008c\u008d\3\2\2\2\u008d\u008e\b\6\1\2\u008e\u00a3\3\2\2\2\u008f"+
		"\u0090\7\b\2\2\u0090\u0091\7\"\2\2\u0091\u0092\7(\2\2\u0092\u0094\5\66"+
		"\34\2\u0093\u0095\7<\2\2\u0094\u0093\3\2\2\2\u0094\u0095\3\2\2\2\u0095"+
		"\u0096\3\2\2\2\u0096\u0097\b\6\1\2\u0097\u00a3\3\2\2\2\u0098\u0099\7\b"+
		"\2\2\u0099\u009a\7\"\2\2\u009a\u009b\78\2\2\u009b\u009c\5<\37\2\u009c"+
		"\u009e\79\2\2\u009d\u009f\7<\2\2\u009e\u009d\3\2\2\2\u009e\u009f\3\2\2"+
		"\2\u009f\u00a0\3\2\2\2\u00a0\u00a1\b\6\1\2\u00a1\u00a3\3\2\2\2\u00a2\u0084"+
		"\3\2\2\2\u00a2\u008f\3\2\2\2\u00a2\u0098\3\2\2\2\u00a3\13\3\2\2\2\u00a4"+
		"\u00a5\7\20\2\2\u00a5\u00a6\7\"\2\2\u00a6\u00a7\78\2\2\u00a7\u00a8\5<"+
		"\37\2\u00a8\u00a9\7(\2\2\u00a9\u00ab\5\66\34\2\u00aa\u00ac\7<\2\2\u00ab"+
		"\u00aa\3\2\2\2\u00ab\u00ac\3\2\2\2\u00ac\u00ad\3\2\2\2\u00ad\u00ae\b\7"+
		"\1\2\u00ae\u00b9\3\2\2\2\u00af\u00b0\7\20\2\2\u00b0\u00b1\7\"\2\2\u00b1"+
		"\u00b2\7(\2\2\u00b2\u00b4\5\66\34\2\u00b3\u00b5\7<\2\2\u00b4\u00b3\3\2"+
		"\2\2\u00b4\u00b5\3\2\2\2\u00b5\u00b6\3\2\2\2\u00b6\u00b7\b\7\1\2\u00b7"+
		"\u00b9\3\2\2\2\u00b8\u00a4\3\2\2\2\u00b8\u00af\3\2\2\2\u00b9\r\3\2\2\2"+
		"\u00ba\u00bc\5\20\t\2\u00bb\u00ba\3\2\2\2\u00bc\u00bd\3\2\2\2\u00bd\u00bb"+
		"\3\2\2\2\u00bd\u00be\3\2\2\2\u00be\u00bf\3\2\2\2\u00bf\u00c0\b\b\1\2\u00c0"+
		"\17\3\2\2\2\u00c1\u00c2\7\r\2\2\u00c2\u00c3\5\66\34\2\u00c3\u00c4\7\66"+
		"\2\2\u00c4\u00c5\5\4\3\2\u00c5\u00c6\7\67\2\2\u00c6\u00c7\b\t\1\2\u00c7"+
		"\u00dd\3\2\2\2\u00c8\u00c9\7\r\2\2\u00c9\u00ca\5\66\34\2\u00ca\u00cb\7"+
		"\66\2\2\u00cb\u00cc\5\4\3\2\u00cc\u00cd\7\67\2\2\u00cd\u00ce\7\16\2\2"+
		"\u00ce\u00cf\7\66\2\2\u00cf\u00d0\5\4\3\2\u00d0\u00d1\7\67\2\2\u00d1\u00d2"+
		"\b\t\1\2\u00d2\u00dd\3\2\2\2\u00d3\u00d4\7\r\2\2\u00d4\u00d5\5\66\34\2"+
		"\u00d5\u00d6\7\66\2\2\u00d6\u00d7\5\4\3\2\u00d7\u00d8\7\67\2\2\u00d8\u00d9"+
		"\7\16\2\2\u00d9\u00da\5\16\b\2\u00da\u00db\b\t\1\2\u00db\u00dd\3\2\2\2"+
		"\u00dc\u00c1\3\2\2\2\u00dc\u00c8\3\2\2\2\u00dc\u00d3\3\2\2\2\u00dd\21"+
		"\3\2\2\2\u00de\u00df\7\23\2\2\u00df\u00e0\5\66\34\2\u00e0\u00e1\7\66\2"+
		"\2\u00e1\u00e2\7\25\2\2\u00e2\u00e3\78\2\2\u00e3\u00e4\5\4\3\2\u00e4\u00e5"+
		"\7\67\2\2\u00e5\u00e6\b\n\1\2\u00e6\u00ef\3\2\2\2\u00e7\u00e8\7\23\2\2"+
		"\u00e8\u00e9\5\66\34\2\u00e9\u00ea\7\66\2\2\u00ea\u00eb\5\24\13\2\u00eb"+
		"\u00ec\7\67\2\2\u00ec\u00ed\b\n\1\2\u00ed\u00ef\3\2\2\2\u00ee\u00de\3"+
		"\2\2\2\u00ee\u00e7\3\2\2\2\u00ef\23\3\2\2\2\u00f0\u00f1\7\24\2\2\u00f1"+
		"\u00f2\5\66\34\2\u00f2\u00f3\78\2\2\u00f3\u00f4\5\4\3\2\u00f4\u00f5\5"+
		"\26\f\2\u00f5\u00f6\b\13\1\2\u00f6\u0101\3\2\2\2\u00f7\u00f8\7\24\2\2"+
		"\u00f8\u00f9\5\66\34\2\u00f9\u00fa\78\2\2\u00fa\u00fb\5\4\3\2\u00fb\u00fc"+
		"\7\25\2\2\u00fc\u00fd\78\2\2\u00fd\u00fe\5\4\3\2\u00fe\u00ff\b\13\1\2"+
		"\u00ff\u0101\3\2\2\2\u0100\u00f0\3\2\2\2\u0100\u00f7\3\2\2\2\u0101\25"+
		"\3\2\2\2\u0102\u0104\5\24\13\2\u0103\u0102\3\2\2\2\u0104\u0105\3\2\2\2"+
		"\u0105\u0103\3\2\2\2\u0105\u0106\3\2\2\2\u0106\u0107\3\2\2\2\u0107\u0108"+
		"\b\f\1\2\u0108\27\3\2\2\2\u0109\u010a\7\"\2\2\u010a\u010b\7(\2\2\u010b"+
		"\u010d\5\66\34\2\u010c\u010e\7<\2\2\u010d\u010c\3\2\2\2\u010d\u010e\3"+
		"\2\2\2\u010e\u010f\3\2\2\2\u010f\u0110\b\r\1\2\u0110\31\3\2\2\2\u0111"+
		"\u0112\7\17\2\2\u0112\u0113\5\66\34\2\u0113\u0114\7\66\2\2\u0114\u0115"+
		"\5\4\3\2\u0115\u0116\7\67\2\2\u0116\u0117\b\16\1\2\u0117\33\3\2\2\2\u0118"+
		"\u0119\7\21\2\2\u0119\u011a\7\"\2\2\u011a\u011b\7\22\2\2\u011b\u011c\5"+
		"\66\34\2\u011c\u011d\7\66\2\2\u011d\u011e\5\4\3\2\u011e\u011f\7\67\2\2"+
		"\u011f\u0120\b\17\1\2\u0120\u012e\3\2\2\2\u0121\u0122\7\21\2\2\u0122\u0123"+
		"\7\"\2\2\u0123\u0124\7\22\2\2\u0124\u0125\5\66\34\2\u0125\u0126\7:\2\2"+
		"\u0126\u0127\7:\2\2\u0127\u0128\5\66\34\2\u0128\u0129\7\66\2\2\u0129\u012a"+
		"\5\4\3\2\u012a\u012b\7\67\2\2\u012b\u012c\b\17\1\2\u012c\u012e\3\2\2\2"+
		"\u012d\u0118\3\2\2\2\u012d\u0121\3\2\2\2\u012e\35\3\2\2\2\u012f\u0130"+
		"\7\26\2\2\u0130\u0131\5\66\34\2\u0131\u0132\7\16\2\2\u0132\u0133\7\66"+
		"\2\2\u0133\u0134\5\4\3\2\u0134\u0135\t\2\2\2\u0135\u0136\7\67\2\2\u0136"+
		"\u0137\b\20\1\2\u0137\37\3\2\2\2\u0138\u013a\7\30\2\2\u0139\u013b\7<\2"+
		"\2\u013a\u0139\3\2\2\2\u013a\u013b\3\2\2\2\u013b\u013c\3\2\2\2\u013c\u014f"+
		"\b\21\1\2\u013d\u013e\7\30\2\2\u013e\u0140\5\66\34\2\u013f\u0141\7<\2"+
		"\2\u0140\u013f\3\2\2\2\u0140\u0141\3\2\2\2\u0141\u0142\3\2\2\2\u0142\u0143"+
		"\b\21\1\2\u0143\u014f\3\2\2\2\u0144\u0146\7\27\2\2\u0145\u0147\7<\2\2"+
		"\u0146\u0145\3\2\2\2\u0146\u0147\3\2\2\2\u0147\u0148\3\2\2\2\u0148\u014f"+
		"\b\21\1\2\u0149\u014b\7\31\2\2\u014a\u014c\7<\2\2\u014b\u014a\3\2\2\2"+
		"\u014b\u014c\3\2\2\2\u014c\u014d\3\2\2\2\u014d\u014f\b\21\1\2\u014e\u0138"+
		"\3\2\2\2\u014e\u013d\3\2\2\2\u014e\u0144\3\2\2\2\u014e\u0149\3\2\2\2\u014f"+
		"!\3\2\2\2\u0150\u0151\7\b\2\2\u0151\u0152\7\"\2\2\u0152\u0153\78\2\2\u0153"+
		"\u0154\7=\2\2\u0154\u0155\5<\37\2\u0155\u0156\7>\2\2\u0156\u0157\7(\2"+
		"\2\u0157\u0158\7=\2\2\u0158\u0159\5\64\33\2\u0159\u015b\7>\2\2\u015a\u015c"+
		"\7<\2\2\u015b\u015a\3\2\2\2\u015b\u015c\3\2\2\2\u015c\u015d\3\2\2\2\u015d"+
		"\u015e\b\22\1\2\u015e\u0186\3\2\2\2\u015f\u0160\7\b\2\2\u0160\u0161\7"+
		"\"\2\2\u0161\u0162\78\2\2\u0162\u0163\7=\2\2\u0163\u0164\5<\37\2\u0164"+
		"\u0165\7>\2\2\u0165\u0166\7(\2\2\u0166\u0167\7=\2\2\u0167\u0169\7>\2\2"+
		"\u0168\u016a\7<\2\2\u0169\u0168\3\2\2\2\u0169\u016a\3\2\2\2\u016a\u016b"+
		"\3\2\2\2\u016b\u016c\b\22\1\2\u016c\u0186\3\2\2\2\u016d\u016e\7\b\2\2"+
		"\u016e\u016f\7\"\2\2\u016f\u0170\7(\2\2\u0170\u0171\7=\2\2\u0171\u0172"+
		"\5<\37\2\u0172\u0173\7>\2\2\u0173\u0174\7\64\2\2\u0174\u0176\7\65\2\2"+
		"\u0175\u0177\7<\2\2\u0176\u0175\3\2\2\2\u0176\u0177\3\2\2\2\u0177\u0186"+
		"\3\2\2\2\u0178\u0179\7\b\2\2\u0179\u017a\7\"\2\2\u017a\u017b\78\2\2\u017b"+
		"\u017c\7=\2\2\u017c\u017d\5<\37\2\u017d\u017e\7>\2\2\u017e\u017f\7(\2"+
		"\2\u017f\u0181\7\"\2\2\u0180\u0182\7<\2\2\u0181\u0180\3\2\2\2\u0181\u0182"+
		"\3\2\2\2\u0182\u0183\3\2\2\2\u0183\u0184\b\22\1\2\u0184\u0186\3\2\2\2"+
		"\u0185\u0150\3\2\2\2\u0185\u015f\3\2\2\2\u0185\u016d\3\2\2\2\u0185\u0178"+
		"\3\2\2\2\u0186#\3\2\2\2\u0187\u0188\7\"\2\2\u0188\u0189\7:\2\2\u0189\u018a"+
		"\7\32\2\2\u018a\u018b\7\64\2\2\u018b\u018c\5\66\34\2\u018c\u018e\7\65"+
		"\2\2\u018d\u018f\7<\2\2\u018e\u018d\3\2\2\2\u018e\u018f\3\2\2\2\u018f"+
		"\u0190\3\2\2\2\u0190\u0191\b\23\1\2\u0191%\3\2\2\2\u0192\u0193\7\"\2\2"+
		"\u0193\u0194\7:\2\2\u0194\u0195\7\33\2\2\u0195\u0196\7\64\2\2\u0196\u0198"+
		"\7\65\2\2\u0197\u0199\7<\2\2\u0198\u0197\3\2\2\2\u0198\u0199\3\2\2\2\u0199"+
		"\u019a\3\2\2\2\u019a\u019b\b\24\1\2\u019b\'\3\2\2\2\u019c\u019d\7\"\2"+
		"\2\u019d\u019e\7:\2\2\u019e\u019f\7\34\2\2\u019f\u01a0\7\64\2\2\u01a0"+
		"\u01a1\7\35\2\2\u01a1\u01a2\78\2\2\u01a2\u01a3\5\66\34\2\u01a3\u01a5\7"+
		"\65\2\2\u01a4\u01a6\7<\2\2\u01a5\u01a4\3\2\2\2\u01a5\u01a6\3\2\2\2\u01a6"+
		"\u01a7\3\2\2\2\u01a7\u01a8\b\25\1\2\u01a8)\3\2\2\2\u01a9\u01aa\7\"\2\2"+
		"\u01aa\u01ab\7:\2\2\u01ab\u01ac\7\36\2\2\u01ac\u01ad\b\26\1\2\u01ad+\3"+
		"\2\2\2\u01ae\u01af\7\"\2\2\u01af\u01b0\7:\2\2\u01b0\u01b1\7\37\2\2\u01b1"+
		"\u01b2\b\27\1\2\u01b2-\3\2\2\2\u01b3\u01b4\7\"\2\2\u01b4\u01b5\7=\2\2"+
		"\u01b5\u01b6\5\66\34\2\u01b6\u01b7\7>\2\2\u01b7\u01b8\b\30\1\2\u01b8/"+
		"\3\2\2\2\u01b9\u01ba\7\"\2\2\u01ba\u01bb\7=\2\2\u01bb\u01bc\5\66\34\2"+
		"\u01bc\u01bd\7>\2\2\u01bd\u01be\7(\2\2\u01be\u01c0\5\66\34\2\u01bf\u01c1"+
		"\7<\2\2\u01c0\u01bf\3\2\2\2\u01c0\u01c1\3\2\2\2\u01c1\u01c2\3\2\2\2\u01c2"+
		"\u01c3\b\31\1\2\u01c3\61\3\2\2\2\u01c4\u01c5\7\b\2\2\u01c5\u01c6\7\"\2"+
		"\2\u01c6\u01c7\7\64\2\2\u01c7\u01c8\78\2\2\u01c8\u01c9\5<\37\2\u01c9\u01ca"+
		"\7\65\2\2\u01ca\u01cc\7(\2\2\u01cb\u01cd\7<\2\2\u01cc\u01cb\3\2\2\2\u01cc"+
		"\u01cd\3\2\2\2\u01cd\63\3\2\2\2\u01ce\u01cf\b\33\1\2\u01cf\u01d0\5\66"+
		"\34\2\u01d0\u01d1\b\33\1\2\u01d1\u01d9\3\2\2\2\u01d2\u01d3\f\4\2\2\u01d3"+
		"\u01d4\7;\2\2\u01d4\u01d5\5\66\34\2\u01d5\u01d6\b\33\1\2\u01d6\u01d8\3"+
		"\2\2\2\u01d7\u01d2\3\2\2\2\u01d8\u01db\3\2\2\2\u01d9\u01d7\3\2\2\2\u01d9"+
		"\u01da\3\2\2\2\u01da\65\3\2\2\2\u01db\u01d9\3\2\2\2\u01dc\u01dd\b\34\1"+
		"\2\u01dd\u01de\7%\2\2\u01de\u01df\5\66\34\16\u01df\u01e0\b\34\1\2\u01e0"+
		"\u0201\3\2\2\2\u01e1\u01e2\7\62\2\2\u01e2\u01e3\5\66\34\r\u01e3\u01e4"+
		"\b\34\1\2\u01e4\u0201\3\2\2\2\u01e5\u01e6\7\64\2\2\u01e6\u01e7\5\66\34"+
		"\2\u01e7\u01e8\7\65\2\2\u01e8\u01e9\b\34\1\2\u01e9\u0201\3\2\2\2\u01ea"+
		"\u01eb\7 \2\2\u01eb\u0201\b\34\1\2\u01ec\u01ed\7!\2\2\u01ed\u0201\b\34"+
		"\1\2\u01ee\u01ef\7\n\2\2\u01ef\u0201\b\34\1\2\u01f0\u01f1\7\13\2\2\u01f1"+
		"\u0201\b\34\1\2\u01f2\u01f3\7\t\2\2\u01f3\u0201\b\34\1\2\u01f4\u01f5\5"+
		"8\35\2\u01f5\u01f6\b\34\1\2\u01f6\u0201\3\2\2\2\u01f7\u01f8\5*\26\2\u01f8"+
		"\u01f9\b\34\1\2\u01f9\u0201\3\2\2\2\u01fa\u01fb\5,\27\2\u01fb\u01fc\b"+
		"\34\1\2\u01fc\u0201\3\2\2\2\u01fd\u01fe\5.\30\2\u01fe\u01ff\b\34\1\2\u01ff"+
		"\u0201\3\2\2\2\u0200\u01dc\3\2\2\2\u0200\u01e1\3\2\2\2\u0200\u01e5\3\2"+
		"\2\2\u0200\u01ea\3\2\2\2\u0200\u01ec\3\2\2\2\u0200\u01ee\3\2\2\2\u0200"+
		"\u01f0\3\2\2\2\u0200\u01f2\3\2\2\2\u0200\u01f4\3\2\2\2\u0200\u01f7\3\2"+
		"\2\2\u0200\u01fa\3\2\2\2\u0200\u01fd\3\2\2\2\u0201\u0227\3\2\2\2\u0202"+
		"\u0203\f\25\2\2\u0203\u0204\t\3\2\2\u0204\u0205\5\66\34\26\u0205\u0206"+
		"\b\34\1\2\u0206\u0226\3\2\2\2\u0207\u0208\f\24\2\2\u0208\u0209\t\4\2\2"+
		"\u0209\u020a\5\66\34\25\u020a\u020b\b\34\1\2\u020b\u0226\3\2\2\2\u020c"+
		"\u020d\f\23\2\2\u020d\u020e\t\5\2\2\u020e\u020f\5\66\34\24\u020f\u0210"+
		"\b\34\1\2\u0210\u0226\3\2\2\2\u0211\u0212\f\22\2\2\u0212\u0213\t\6\2\2"+
		"\u0213\u0214\5\66\34\23\u0214\u0215\b\34\1\2\u0215\u0226\3\2\2\2\u0216"+
		"\u0217\f\21\2\2\u0217\u0218\t\7\2\2\u0218\u0219\5\66\34\22\u0219\u021a"+
		"\b\34\1\2\u021a\u0226\3\2\2\2\u021b\u021c\f\20\2\2\u021c\u021d\7\'\2\2"+
		"\u021d\u021e\5\66\34\21\u021e\u021f\b\34\1\2\u021f\u0226\3\2\2\2\u0220"+
		"\u0221\f\17\2\2\u0221\u0222\7&\2\2\u0222\u0223\5\66\34\20\u0223\u0224"+
		"\b\34\1\2\u0224\u0226\3\2\2\2\u0225\u0202\3\2\2\2\u0225\u0207\3\2\2\2"+
		"\u0225\u020c\3\2\2\2\u0225\u0211\3\2\2\2\u0225\u0216\3\2\2\2\u0225\u021b"+
		"\3\2\2\2\u0225\u0220\3\2\2\2\u0226\u0229\3\2\2\2\u0227\u0225\3\2\2\2\u0227"+
		"\u0228\3\2\2\2\u0228\67\3\2\2\2\u0229\u0227\3\2\2\2\u022a\u022b\7\"\2"+
		"\2\u022b\u022c\b\35\1\2\u022c9\3\2\2\2\u022d\u022e\7\"\2\2\u022e\u022f"+
		"\7)\2\2\u022f\u0231\5\66\34\2\u0230\u0232\7<\2\2\u0231\u0230\3\2\2\2\u0231"+
		"\u0232\3\2\2\2\u0232\u0233\3\2\2\2\u0233\u0234\b\36\1\2\u0234\u023e\3"+
		"\2\2\2\u0235\u0236\7\"\2\2\u0236\u0237\7*\2\2\u0237\u0239\5\66\34\2\u0238"+
		"\u023a\7<\2\2\u0239\u0238\3\2\2\2\u0239\u023a\3\2\2\2\u023a\u023b\3\2"+
		"\2\2\u023b\u023c\b\36\1\2\u023c\u023e\3\2\2\2\u023d\u022d\3\2\2\2\u023d"+
		"\u0235\3\2\2\2\u023e;\3\2\2\2\u023f\u0240\7\3\2\2\u0240\u024a\b\37\1\2"+
		"\u0241\u0242\7\4\2\2\u0242\u024a\b\37\1\2\u0243\u0244\7\5\2\2\u0244\u024a"+
		"\b\37\1\2\u0245\u0246\7\6\2\2\u0246\u024a\b\37\1\2\u0247\u0248\7\7\2\2"+
		"\u0248\u024a\b\37\1\2\u0249\u023f\3\2\2\2\u0249\u0241\3\2\2\2\u0249\u0243"+
		"\3\2\2\2\u0249\u0245\3\2\2\2\u0249\u0247\3\2\2\2\u024a=\3\2\2\2*Ey\u0080"+
		"\u008b\u0094\u009e\u00a2\u00ab\u00b4\u00b8\u00bd\u00dc\u00ee\u0100\u0105"+
		"\u010d\u012d\u013a\u0140\u0146\u014b\u014e\u015b\u0169\u0176\u0181\u0185"+
		"\u018e\u0198\u01a5\u01c0\u01cc\u01d9\u0200\u0225\u0227\u0231\u0239\u023d"+
		"\u0249";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
// Generated from c:\Users\50246\Desktop\Ingenierķa-Sexto Semestre\Compiladores 2\Proyecto 1\Server\Lex.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class Lex extends Lexer {
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
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"STRINGS", "INTS", "FLOATS", "BOOLS", "CHARACTERS", "VAR", "NIL", "TRUE", 
			"FALSE", "PRINT", "IF", "ELSE", "WHILE", "LET", "FOR", "IN", "SWITCH", 
			"CASE", "DEFAULT", "GUARD", "CONTINUE", "RETURN", "BREAK", "APPEND", 
			"REMOVELAST", "REMOVE", "AT", "ISEMPTY", "COUNT", "NUMBER", "STRING", 
			"ID", "DIF", "IG_IG", "NOT", "OR", "AND", "IG", "IG_ADD", "IG_SUB", "MAY_IG", 
			"MEN_IG", "MAYOR", "MENOR", "MUL", "DIV", "ADD", "SUB", "PERCENT", "PARIZQ", 
			"PARDER", "LLAVEIZQ", "LLAVEDER", "DOUBLEPTS", "QUESTION", "POINT", "COMA", 
			"PTCOMA", "CORCHETEIZQ", "CORCHETEDER", "WHITESPACE", "COMMENT", "LINE_COMMENT", 
			"ESC_SEQ"
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


	public Lex(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Lex.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2A\u01af\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6"+
		"\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\n\3"+
		"\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\r\3\r\3"+
		"\r\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3\20\3\20"+
		"\3\20\3\20\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\23\3\23"+
		"\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\25\3\25\3\25"+
		"\3\25\3\25\3\25\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\27\3\27"+
		"\3\27\3\27\3\27\3\27\3\27\3\30\3\30\3\30\3\30\3\30\3\30\3\31\3\31\3\31"+
		"\3\31\3\31\3\31\3\31\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32"+
		"\3\32\3\33\3\33\3\33\3\33\3\33\3\33\3\33\3\34\3\34\3\34\3\35\3\35\3\35"+
		"\3\35\3\35\3\35\3\35\3\35\3\36\3\36\3\36\3\36\3\36\3\36\3\37\6\37\u0131"+
		"\n\37\r\37\16\37\u0132\3\37\3\37\6\37\u0137\n\37\r\37\16\37\u0138\5\37"+
		"\u013b\n\37\3 \3 \7 \u013f\n \f \16 \u0142\13 \3 \3 \3!\3!\7!\u0148\n"+
		"!\f!\16!\u014b\13!\3\"\3\"\3\"\3#\3#\3#\3$\3$\3%\3%\3%\3&\3&\3&\3\'\3"+
		"\'\3(\3(\3(\3)\3)\3)\3*\3*\3*\3+\3+\3+\3,\3,\3-\3-\3.\3.\3/\3/\3\60\3"+
		"\60\3\61\3\61\3\62\3\62\3\63\3\63\3\64\3\64\3\65\3\65\3\66\3\66\3\67\3"+
		"\67\38\38\39\39\3:\3:\3;\3;\3<\3<\3=\3=\3>\6>\u018e\n>\r>\16>\u018f\3"+
		">\3>\3?\3?\3?\3?\7?\u0198\n?\f?\16?\u019b\13?\3?\3?\3?\3?\3?\3@\3@\3@"+
		"\3@\7@\u01a6\n@\f@\16@\u01a9\13@\3@\3@\3A\3A\3A\3\u0199\2B\3\3\5\4\7\5"+
		"\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23"+
		"%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G"+
		"%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62c\63e\64g\65i\66k\67m8o9q:s;u<w=y>{"+
		"?}@\177A\u0081\2\3\2\t\3\2\62;\3\2$$\4\2C\\c|\6\2\62;C\\aac|\6\2\13\f"+
		"\17\17\"\"^^\4\2\f\f\17\17\t\2\"#%%--/\60<<BB]_\2\u01b5\2\3\3\2\2\2\2"+
		"\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2"+
		"\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2"+
		"\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2"+
		"\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2"+
		"\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2"+
		"\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2"+
		"K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3"+
		"\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2"+
		"\2\2e\3\2\2\2\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2\2m\3\2\2\2\2o\3\2\2\2\2"+
		"q\3\2\2\2\2s\3\2\2\2\2u\3\2\2\2\2w\3\2\2\2\2y\3\2\2\2\2{\3\2\2\2\2}\3"+
		"\2\2\2\2\177\3\2\2\2\3\u0083\3\2\2\2\5\u008a\3\2\2\2\7\u008e\3\2\2\2\t"+
		"\u0094\3\2\2\2\13\u0099\3\2\2\2\r\u00a3\3\2\2\2\17\u00a7\3\2\2\2\21\u00ab"+
		"\3\2\2\2\23\u00b0\3\2\2\2\25\u00b6\3\2\2\2\27\u00bc\3\2\2\2\31\u00bf\3"+
		"\2\2\2\33\u00c4\3\2\2\2\35\u00ca\3\2\2\2\37\u00ce\3\2\2\2!\u00d2\3\2\2"+
		"\2#\u00d5\3\2\2\2%\u00dc\3\2\2\2\'\u00e1\3\2\2\2)\u00e9\3\2\2\2+\u00ef"+
		"\3\2\2\2-\u00f8\3\2\2\2/\u00ff\3\2\2\2\61\u0105\3\2\2\2\63\u010c\3\2\2"+
		"\2\65\u0117\3\2\2\2\67\u011e\3\2\2\29\u0121\3\2\2\2;\u0129\3\2\2\2=\u0130"+
		"\3\2\2\2?\u013c\3\2\2\2A\u0145\3\2\2\2C\u014c\3\2\2\2E\u014f\3\2\2\2G"+
		"\u0152\3\2\2\2I\u0154\3\2\2\2K\u0157\3\2\2\2M\u015a\3\2\2\2O\u015c\3\2"+
		"\2\2Q\u015f\3\2\2\2S\u0162\3\2\2\2U\u0165\3\2\2\2W\u0168\3\2\2\2Y\u016a"+
		"\3\2\2\2[\u016c\3\2\2\2]\u016e\3\2\2\2_\u0170\3\2\2\2a\u0172\3\2\2\2c"+
		"\u0174\3\2\2\2e\u0176\3\2\2\2g\u0178\3\2\2\2i\u017a\3\2\2\2k\u017c\3\2"+
		"\2\2m\u017e\3\2\2\2o\u0180\3\2\2\2q\u0182\3\2\2\2s\u0184\3\2\2\2u\u0186"+
		"\3\2\2\2w\u0188\3\2\2\2y\u018a\3\2\2\2{\u018d\3\2\2\2}\u0193\3\2\2\2\177"+
		"\u01a1\3\2\2\2\u0081\u01ac\3\2\2\2\u0083\u0084\7U\2\2\u0084\u0085\7v\2"+
		"\2\u0085\u0086\7t\2\2\u0086\u0087\7k\2\2\u0087\u0088\7p\2\2\u0088\u0089"+
		"\7i\2\2\u0089\4\3\2\2\2\u008a\u008b\7K\2\2\u008b\u008c\7p\2\2\u008c\u008d"+
		"\7v\2\2\u008d\6\3\2\2\2\u008e\u008f\7H\2\2\u008f\u0090\7n\2\2\u0090\u0091"+
		"\7q\2\2\u0091\u0092\7c\2\2\u0092\u0093\7v\2\2\u0093\b\3\2\2\2\u0094\u0095"+
		"\7D\2\2\u0095\u0096\7q\2\2\u0096\u0097\7q\2\2\u0097\u0098\7n\2\2\u0098"+
		"\n\3\2\2\2\u0099\u009a\7E\2\2\u009a\u009b\7j\2\2\u009b\u009c\7c\2\2\u009c"+
		"\u009d\7t\2\2\u009d\u009e\7c\2\2\u009e\u009f\7e\2\2\u009f\u00a0\7v\2\2"+
		"\u00a0\u00a1\7g\2\2\u00a1\u00a2\7t\2\2\u00a2\f\3\2\2\2\u00a3\u00a4\7x"+
		"\2\2\u00a4\u00a5\7c\2\2\u00a5\u00a6\7t\2\2\u00a6\16\3\2\2\2\u00a7\u00a8"+
		"\7p\2\2\u00a8\u00a9\7k\2\2\u00a9\u00aa\7n\2\2\u00aa\20\3\2\2\2\u00ab\u00ac"+
		"\7v\2\2\u00ac\u00ad\7t\2\2\u00ad\u00ae\7w\2\2\u00ae\u00af\7g\2\2\u00af"+
		"\22\3\2\2\2\u00b0\u00b1\7h\2\2\u00b1\u00b2\7c\2\2\u00b2\u00b3\7n\2\2\u00b3"+
		"\u00b4\7u\2\2\u00b4\u00b5\7g\2\2\u00b5\24\3\2\2\2\u00b6\u00b7\7r\2\2\u00b7"+
		"\u00b8\7t\2\2\u00b8\u00b9\7k\2\2\u00b9\u00ba\7p\2\2\u00ba\u00bb\7v\2\2"+
		"\u00bb\26\3\2\2\2\u00bc\u00bd\7k\2\2\u00bd\u00be\7h\2\2\u00be\30\3\2\2"+
		"\2\u00bf\u00c0\7g\2\2\u00c0\u00c1\7n\2\2\u00c1\u00c2\7u\2\2\u00c2\u00c3"+
		"\7g\2\2\u00c3\32\3\2\2\2\u00c4\u00c5\7y\2\2\u00c5\u00c6\7j\2\2\u00c6\u00c7"+
		"\7k\2\2\u00c7\u00c8\7n\2\2\u00c8\u00c9\7g\2\2\u00c9\34\3\2\2\2\u00ca\u00cb"+
		"\7n\2\2\u00cb\u00cc\7g\2\2\u00cc\u00cd\7v\2\2\u00cd\36\3\2\2\2\u00ce\u00cf"+
		"\7h\2\2\u00cf\u00d0\7q\2\2\u00d0\u00d1\7t\2\2\u00d1 \3\2\2\2\u00d2\u00d3"+
		"\7k\2\2\u00d3\u00d4\7p\2\2\u00d4\"\3\2\2\2\u00d5\u00d6\7u\2\2\u00d6\u00d7"+
		"\7y\2\2\u00d7\u00d8\7k\2\2\u00d8\u00d9\7v\2\2\u00d9\u00da\7e\2\2\u00da"+
		"\u00db\7j\2\2\u00db$\3\2\2\2\u00dc\u00dd\7e\2\2\u00dd\u00de\7c\2\2\u00de"+
		"\u00df\7u\2\2\u00df\u00e0\7g\2\2\u00e0&\3\2\2\2\u00e1\u00e2\7f\2\2\u00e2"+
		"\u00e3\7g\2\2\u00e3\u00e4\7h\2\2\u00e4\u00e5\7c\2\2\u00e5\u00e6\7w\2\2"+
		"\u00e6\u00e7\7n\2\2\u00e7\u00e8\7v\2\2\u00e8(\3\2\2\2\u00e9\u00ea\7i\2"+
		"\2\u00ea\u00eb\7w\2\2\u00eb\u00ec\7c\2\2\u00ec\u00ed\7t\2\2\u00ed\u00ee"+
		"\7f\2\2\u00ee*\3\2\2\2\u00ef\u00f0\7e\2\2\u00f0\u00f1\7q\2\2\u00f1\u00f2"+
		"\7p\2\2\u00f2\u00f3\7v\2\2\u00f3\u00f4\7k\2\2\u00f4\u00f5\7p\2\2\u00f5"+
		"\u00f6\7w\2\2\u00f6\u00f7\7g\2\2\u00f7,\3\2\2\2\u00f8\u00f9\7t\2\2\u00f9"+
		"\u00fa\7g\2\2\u00fa\u00fb\7v\2\2\u00fb\u00fc\7w\2\2\u00fc\u00fd\7t\2\2"+
		"\u00fd\u00fe\7p\2\2\u00fe.\3\2\2\2\u00ff\u0100\7d\2\2\u0100\u0101\7t\2"+
		"\2\u0101\u0102\7g\2\2\u0102\u0103\7c\2\2\u0103\u0104\7m\2\2\u0104\60\3"+
		"\2\2\2\u0105\u0106\7c\2\2\u0106\u0107\7r\2\2\u0107\u0108\7r\2\2\u0108"+
		"\u0109\7g\2\2\u0109\u010a\7p\2\2\u010a\u010b\7f\2\2\u010b\62\3\2\2\2\u010c"+
		"\u010d\7t\2\2\u010d\u010e\7g\2\2\u010e\u010f\7o\2\2\u010f\u0110\7q\2\2"+
		"\u0110\u0111\7x\2\2\u0111\u0112\7g\2\2\u0112\u0113\7N\2\2\u0113\u0114"+
		"\7c\2\2\u0114\u0115\7u\2\2\u0115\u0116\7v\2\2\u0116\64\3\2\2\2\u0117\u0118"+
		"\7t\2\2\u0118\u0119\7g\2\2\u0119\u011a\7o\2\2\u011a\u011b\7q\2\2\u011b"+
		"\u011c\7x\2\2\u011c\u011d\7g\2\2\u011d\66\3\2\2\2\u011e\u011f\7c\2\2\u011f"+
		"\u0120\7v\2\2\u01208\3\2\2\2\u0121\u0122\7K\2\2\u0122\u0123\7u\2\2\u0123"+
		"\u0124\7G\2\2\u0124\u0125\7o\2\2\u0125\u0126\7r\2\2\u0126\u0127\7v\2\2"+
		"\u0127\u0128\7{\2\2\u0128:\3\2\2\2\u0129\u012a\7e\2\2\u012a\u012b\7q\2"+
		"\2\u012b\u012c\7w\2\2\u012c\u012d\7p\2\2\u012d\u012e\7v\2\2\u012e<\3\2"+
		"\2\2\u012f\u0131\t\2\2\2\u0130\u012f\3\2\2\2\u0131\u0132\3\2\2\2\u0132"+
		"\u0130\3\2\2\2\u0132\u0133\3\2\2\2\u0133\u013a\3\2\2\2\u0134\u0136\7\60"+
		"\2\2\u0135\u0137\t\2\2\2\u0136\u0135\3\2\2\2\u0137\u0138\3\2\2\2\u0138"+
		"\u0136\3\2\2\2\u0138\u0139\3\2\2\2\u0139\u013b\3\2\2\2\u013a\u0134\3\2"+
		"\2\2\u013a\u013b\3\2\2\2\u013b>\3\2\2\2\u013c\u0140\7$\2\2\u013d\u013f"+
		"\n\3\2\2\u013e\u013d\3\2\2\2\u013f\u0142\3\2\2\2\u0140\u013e\3\2\2\2\u0140"+
		"\u0141\3\2\2\2\u0141\u0143\3\2\2\2\u0142\u0140\3\2\2\2\u0143\u0144\7$"+
		"\2\2\u0144@\3\2\2\2\u0145\u0149\t\4\2\2\u0146\u0148\t\5\2\2\u0147\u0146"+
		"\3\2\2\2\u0148\u014b\3\2\2\2\u0149\u0147\3\2\2\2\u0149\u014a\3\2\2\2\u014a"+
		"B\3\2\2\2\u014b\u0149\3\2\2\2\u014c\u014d\7#\2\2\u014d\u014e\7?\2\2\u014e"+
		"D\3\2\2\2\u014f\u0150\7?\2\2\u0150\u0151\7?\2\2\u0151F\3\2\2\2\u0152\u0153"+
		"\7#\2\2\u0153H\3\2\2\2\u0154\u0155\7~\2\2\u0155\u0156\7~\2\2\u0156J\3"+
		"\2\2\2\u0157\u0158\7(\2\2\u0158\u0159\7(\2\2\u0159L\3\2\2\2\u015a\u015b"+
		"\7?\2\2\u015bN\3\2\2\2\u015c\u015d\7-\2\2\u015d\u015e\7?\2\2\u015eP\3"+
		"\2\2\2\u015f\u0160\7/\2\2\u0160\u0161\7?\2\2\u0161R\3\2\2\2\u0162\u0163"+
		"\7@\2\2\u0163\u0164\7?\2\2\u0164T\3\2\2\2\u0165\u0166\7>\2\2\u0166\u0167"+
		"\7?\2\2\u0167V\3\2\2\2\u0168\u0169\7@\2\2\u0169X\3\2\2\2\u016a\u016b\7"+
		">\2\2\u016bZ\3\2\2\2\u016c\u016d\7,\2\2\u016d\\\3\2\2\2\u016e\u016f\7"+
		"\61\2\2\u016f^\3\2\2\2\u0170\u0171\7-\2\2\u0171`\3\2\2\2\u0172\u0173\7"+
		"/\2\2\u0173b\3\2\2\2\u0174\u0175\7\'\2\2\u0175d\3\2\2\2\u0176\u0177\7"+
		"*\2\2\u0177f\3\2\2\2\u0178\u0179\7+\2\2\u0179h\3\2\2\2\u017a\u017b\7}"+
		"\2\2\u017bj\3\2\2\2\u017c\u017d\7\177\2\2\u017dl\3\2\2\2\u017e\u017f\7"+
		"<\2\2\u017fn\3\2\2\2\u0180\u0181\7A\2\2\u0181p\3\2\2\2\u0182\u0183\7\60"+
		"\2\2\u0183r\3\2\2\2\u0184\u0185\7.\2\2\u0185t\3\2\2\2\u0186\u0187\7=\2"+
		"\2\u0187v\3\2\2\2\u0188\u0189\7]\2\2\u0189x\3\2\2\2\u018a\u018b\7_\2\2"+
		"\u018bz\3\2\2\2\u018c\u018e\t\6\2\2\u018d\u018c\3\2\2\2\u018e\u018f\3"+
		"\2\2\2\u018f\u018d\3\2\2\2\u018f\u0190\3\2\2\2\u0190\u0191\3\2\2\2\u0191"+
		"\u0192\b>\2\2\u0192|\3\2\2\2\u0193\u0194\7\61\2\2\u0194\u0195\7,\2\2\u0195"+
		"\u0199\3\2\2\2\u0196\u0198\13\2\2\2\u0197\u0196\3\2\2\2\u0198\u019b\3"+
		"\2\2\2\u0199\u019a\3\2\2\2\u0199\u0197\3\2\2\2\u019a\u019c\3\2\2\2\u019b"+
		"\u0199\3\2\2\2\u019c\u019d\7,\2\2\u019d\u019e\7\61\2\2\u019e\u019f\3\2"+
		"\2\2\u019f\u01a0\b?\2\2\u01a0~\3\2\2\2\u01a1\u01a2\7\61\2\2\u01a2\u01a3"+
		"\7\61\2\2\u01a3\u01a7\3\2\2\2\u01a4\u01a6\n\7\2\2\u01a5\u01a4\3\2\2\2"+
		"\u01a6\u01a9\3\2\2\2\u01a7\u01a5\3\2\2\2\u01a7\u01a8\3\2\2\2\u01a8\u01aa"+
		"\3\2\2\2\u01a9\u01a7\3\2\2\2\u01aa\u01ab\b@\2\2\u01ab\u0080\3\2\2\2\u01ac"+
		"\u01ad\7^\2\2\u01ad\u01ae\t\b\2\2\u01ae\u0082\3\2\2\2\13\2\u0132\u0138"+
		"\u013a\u0140\u0149\u018f\u0199\u01a7\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
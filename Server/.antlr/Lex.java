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
		FALSE=9, PRINT=10, IF=11, ELSE=12, WHILE=13, LET=14, FOR=15, IN=16, NUMBER=17, 
		STRING=18, ID=19, DIF=20, IG_IG=21, NOT=22, OR=23, AND=24, IG=25, IG_ADD=26, 
		IG_SUB=27, MAY_IG=28, MEN_IG=29, MAYOR=30, MENOR=31, MUL=32, DIV=33, ADD=34, 
		SUB=35, PERCENT=36, PARIZQ=37, PARDER=38, LLAVEIZQ=39, LLAVEDER=40, DOUBLEPTS=41, 
		QUESTION=42, WHITESPACE=43, COMMENT=44, LINE_COMMENT=45;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"STRINGS", "INTS", "FLOATS", "BOOLS", "CHARACTERS", "VAR", "NIL", "TRUE", 
			"FALSE", "PRINT", "IF", "ELSE", "WHILE", "LET", "FOR", "IN", "NUMBER", 
			"STRING", "ID", "DIF", "IG_IG", "NOT", "OR", "AND", "IG", "IG_ADD", "IG_SUB", 
			"MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", "DIV", "ADD", "SUB", "PERCENT", 
			"PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "DOUBLEPTS", "QUESTION", 
			"WHITESPACE", "COMMENT", "LINE_COMMENT", "ESC_SEQ"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'String'", "'Int'", "'Float'", "'Bool'", "'Character'", "'var'", 
			"'nil'", "'true'", "'false'", "'print'", "'if'", "'else'", "'while'", 
			"'let'", "'for'", "'in'", null, null, null, "'!='", "'=='", "'!'", "'||'", 
			"'&&'", "'='", "'+='", "'-='", "'>='", "'<='", "'>'", "'<'", "'*'", "'/'", 
			"'+'", "'-'", "'%'", "'('", "')'", "'{'", "'}'", "':'", "'?'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "STRINGS", "INTS", "FLOATS", "BOOLS", "CHARACTERS", "VAR", "NIL", 
			"TRUE", "FALSE", "PRINT", "IF", "ELSE", "WHILE", "LET", "FOR", "IN", 
			"NUMBER", "STRING", "ID", "DIF", "IG_IG", "NOT", "OR", "AND", "IG", "IG_ADD", 
			"IG_SUB", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", "DIV", "ADD", 
			"SUB", "PERCENT", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "DOUBLEPTS", 
			"QUESTION", "WHITESPACE", "COMMENT", "LINE_COMMENT"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2/\u0127\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\3\6\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\n\3\n"+
		"\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\r\3\r\3\r"+
		"\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3\20\3\20\3"+
		"\20\3\20\3\21\3\21\3\21\3\22\6\22\u00b3\n\22\r\22\16\22\u00b4\3\22\3\22"+
		"\6\22\u00b9\n\22\r\22\16\22\u00ba\5\22\u00bd\n\22\3\23\3\23\7\23\u00c1"+
		"\n\23\f\23\16\23\u00c4\13\23\3\23\3\23\3\24\3\24\7\24\u00ca\n\24\f\24"+
		"\16\24\u00cd\13\24\3\25\3\25\3\25\3\26\3\26\3\26\3\27\3\27\3\30\3\30\3"+
		"\30\3\31\3\31\3\31\3\32\3\32\3\33\3\33\3\33\3\34\3\34\3\34\3\35\3\35\3"+
		"\35\3\36\3\36\3\36\3\37\3\37\3 \3 \3!\3!\3\"\3\"\3#\3#\3$\3$\3%\3%\3&"+
		"\3&\3\'\3\'\3(\3(\3)\3)\3*\3*\3+\3+\3,\6,\u0106\n,\r,\16,\u0107\3,\3,"+
		"\3-\3-\3-\3-\7-\u0110\n-\f-\16-\u0113\13-\3-\3-\3-\3-\3-\3.\3.\3.\3.\7"+
		".\u011e\n.\f.\16.\u0121\13.\3.\3.\3/\3/\3/\3\u0111\2\60\3\3\5\4\7\5\t"+
		"\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23"+
		"%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G"+
		"%I&K\'M(O)Q*S+U,W-Y.[/]\2\3\2\t\3\2\62;\3\2$$\4\2C\\c|\6\2\62;C\\aac|"+
		"\6\2\13\f\17\17\"\"^^\4\2\f\f\17\17\t\2\"#%%--/\60<<BB]_\2\u012d\2\3\3"+
		"\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2"+
		"\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3"+
		"\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2"+
		"%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61"+
		"\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2"+
		"\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I"+
		"\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2"+
		"\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\3_\3\2\2\2\5f\3\2\2\2\7j\3\2\2\2"+
		"\tp\3\2\2\2\13u\3\2\2\2\r\177\3\2\2\2\17\u0083\3\2\2\2\21\u0087\3\2\2"+
		"\2\23\u008c\3\2\2\2\25\u0092\3\2\2\2\27\u0098\3\2\2\2\31\u009b\3\2\2\2"+
		"\33\u00a0\3\2\2\2\35\u00a6\3\2\2\2\37\u00aa\3\2\2\2!\u00ae\3\2\2\2#\u00b2"+
		"\3\2\2\2%\u00be\3\2\2\2\'\u00c7\3\2\2\2)\u00ce\3\2\2\2+\u00d1\3\2\2\2"+
		"-\u00d4\3\2\2\2/\u00d6\3\2\2\2\61\u00d9\3\2\2\2\63\u00dc\3\2\2\2\65\u00de"+
		"\3\2\2\2\67\u00e1\3\2\2\29\u00e4\3\2\2\2;\u00e7\3\2\2\2=\u00ea\3\2\2\2"+
		"?\u00ec\3\2\2\2A\u00ee\3\2\2\2C\u00f0\3\2\2\2E\u00f2\3\2\2\2G\u00f4\3"+
		"\2\2\2I\u00f6\3\2\2\2K\u00f8\3\2\2\2M\u00fa\3\2\2\2O\u00fc\3\2\2\2Q\u00fe"+
		"\3\2\2\2S\u0100\3\2\2\2U\u0102\3\2\2\2W\u0105\3\2\2\2Y\u010b\3\2\2\2["+
		"\u0119\3\2\2\2]\u0124\3\2\2\2_`\7U\2\2`a\7v\2\2ab\7t\2\2bc\7k\2\2cd\7"+
		"p\2\2de\7i\2\2e\4\3\2\2\2fg\7K\2\2gh\7p\2\2hi\7v\2\2i\6\3\2\2\2jk\7H\2"+
		"\2kl\7n\2\2lm\7q\2\2mn\7c\2\2no\7v\2\2o\b\3\2\2\2pq\7D\2\2qr\7q\2\2rs"+
		"\7q\2\2st\7n\2\2t\n\3\2\2\2uv\7E\2\2vw\7j\2\2wx\7c\2\2xy\7t\2\2yz\7c\2"+
		"\2z{\7e\2\2{|\7v\2\2|}\7g\2\2}~\7t\2\2~\f\3\2\2\2\177\u0080\7x\2\2\u0080"+
		"\u0081\7c\2\2\u0081\u0082\7t\2\2\u0082\16\3\2\2\2\u0083\u0084\7p\2\2\u0084"+
		"\u0085\7k\2\2\u0085\u0086\7n\2\2\u0086\20\3\2\2\2\u0087\u0088\7v\2\2\u0088"+
		"\u0089\7t\2\2\u0089\u008a\7w\2\2\u008a\u008b\7g\2\2\u008b\22\3\2\2\2\u008c"+
		"\u008d\7h\2\2\u008d\u008e\7c\2\2\u008e\u008f\7n\2\2\u008f\u0090\7u\2\2"+
		"\u0090\u0091\7g\2\2\u0091\24\3\2\2\2\u0092\u0093\7r\2\2\u0093\u0094\7"+
		"t\2\2\u0094\u0095\7k\2\2\u0095\u0096\7p\2\2\u0096\u0097\7v\2\2\u0097\26"+
		"\3\2\2\2\u0098\u0099\7k\2\2\u0099\u009a\7h\2\2\u009a\30\3\2\2\2\u009b"+
		"\u009c\7g\2\2\u009c\u009d\7n\2\2\u009d\u009e\7u\2\2\u009e\u009f\7g\2\2"+
		"\u009f\32\3\2\2\2\u00a0\u00a1\7y\2\2\u00a1\u00a2\7j\2\2\u00a2\u00a3\7"+
		"k\2\2\u00a3\u00a4\7n\2\2\u00a4\u00a5\7g\2\2\u00a5\34\3\2\2\2\u00a6\u00a7"+
		"\7n\2\2\u00a7\u00a8\7g\2\2\u00a8\u00a9\7v\2\2\u00a9\36\3\2\2\2\u00aa\u00ab"+
		"\7h\2\2\u00ab\u00ac\7q\2\2\u00ac\u00ad\7t\2\2\u00ad \3\2\2\2\u00ae\u00af"+
		"\7k\2\2\u00af\u00b0\7p\2\2\u00b0\"\3\2\2\2\u00b1\u00b3\t\2\2\2\u00b2\u00b1"+
		"\3\2\2\2\u00b3\u00b4\3\2\2\2\u00b4\u00b2\3\2\2\2\u00b4\u00b5\3\2\2\2\u00b5"+
		"\u00bc\3\2\2\2\u00b6\u00b8\7\60\2\2\u00b7\u00b9\t\2\2\2\u00b8\u00b7\3"+
		"\2\2\2\u00b9\u00ba\3\2\2\2\u00ba\u00b8\3\2\2\2\u00ba\u00bb\3\2\2\2\u00bb"+
		"\u00bd\3\2\2\2\u00bc\u00b6\3\2\2\2\u00bc\u00bd\3\2\2\2\u00bd$\3\2\2\2"+
		"\u00be\u00c2\7$\2\2\u00bf\u00c1\n\3\2\2\u00c0\u00bf\3\2\2\2\u00c1\u00c4"+
		"\3\2\2\2\u00c2\u00c0\3\2\2\2\u00c2\u00c3\3\2\2\2\u00c3\u00c5\3\2\2\2\u00c4"+
		"\u00c2\3\2\2\2\u00c5\u00c6\7$\2\2\u00c6&\3\2\2\2\u00c7\u00cb\t\4\2\2\u00c8"+
		"\u00ca\t\5\2\2\u00c9\u00c8\3\2\2\2\u00ca\u00cd\3\2\2\2\u00cb\u00c9\3\2"+
		"\2\2\u00cb\u00cc\3\2\2\2\u00cc(\3\2\2\2\u00cd\u00cb\3\2\2\2\u00ce\u00cf"+
		"\7#\2\2\u00cf\u00d0\7?\2\2\u00d0*\3\2\2\2\u00d1\u00d2\7?\2\2\u00d2\u00d3"+
		"\7?\2\2\u00d3,\3\2\2\2\u00d4\u00d5\7#\2\2\u00d5.\3\2\2\2\u00d6\u00d7\7"+
		"~\2\2\u00d7\u00d8\7~\2\2\u00d8\60\3\2\2\2\u00d9\u00da\7(\2\2\u00da\u00db"+
		"\7(\2\2\u00db\62\3\2\2\2\u00dc\u00dd\7?\2\2\u00dd\64\3\2\2\2\u00de\u00df"+
		"\7-\2\2\u00df\u00e0\7?\2\2\u00e0\66\3\2\2\2\u00e1\u00e2\7/\2\2\u00e2\u00e3"+
		"\7?\2\2\u00e38\3\2\2\2\u00e4\u00e5\7@\2\2\u00e5\u00e6\7?\2\2\u00e6:\3"+
		"\2\2\2\u00e7\u00e8\7>\2\2\u00e8\u00e9\7?\2\2\u00e9<\3\2\2\2\u00ea\u00eb"+
		"\7@\2\2\u00eb>\3\2\2\2\u00ec\u00ed\7>\2\2\u00ed@\3\2\2\2\u00ee\u00ef\7"+
		",\2\2\u00efB\3\2\2\2\u00f0\u00f1\7\61\2\2\u00f1D\3\2\2\2\u00f2\u00f3\7"+
		"-\2\2\u00f3F\3\2\2\2\u00f4\u00f5\7/\2\2\u00f5H\3\2\2\2\u00f6\u00f7\7\'"+
		"\2\2\u00f7J\3\2\2\2\u00f8\u00f9\7*\2\2\u00f9L\3\2\2\2\u00fa\u00fb\7+\2"+
		"\2\u00fbN\3\2\2\2\u00fc\u00fd\7}\2\2\u00fdP\3\2\2\2\u00fe\u00ff\7\177"+
		"\2\2\u00ffR\3\2\2\2\u0100\u0101\7<\2\2\u0101T\3\2\2\2\u0102\u0103\7A\2"+
		"\2\u0103V\3\2\2\2\u0104\u0106\t\6\2\2\u0105\u0104\3\2\2\2\u0106\u0107"+
		"\3\2\2\2\u0107\u0105\3\2\2\2\u0107\u0108\3\2\2\2\u0108\u0109\3\2\2\2\u0109"+
		"\u010a\b,\2\2\u010aX\3\2\2\2\u010b\u010c\7\61\2\2\u010c\u010d\7,\2\2\u010d"+
		"\u0111\3\2\2\2\u010e\u0110\13\2\2\2\u010f\u010e\3\2\2\2\u0110\u0113\3"+
		"\2\2\2\u0111\u0112\3\2\2\2\u0111\u010f\3\2\2\2\u0112\u0114\3\2\2\2\u0113"+
		"\u0111\3\2\2\2\u0114\u0115\7,\2\2\u0115\u0116\7\61\2\2\u0116\u0117\3\2"+
		"\2\2\u0117\u0118\b-\2\2\u0118Z\3\2\2\2\u0119\u011a\7\61\2\2\u011a\u011b"+
		"\7\61\2\2\u011b\u011f\3\2\2\2\u011c\u011e\n\7\2\2\u011d\u011c\3\2\2\2"+
		"\u011e\u0121\3\2\2\2\u011f\u011d\3\2\2\2\u011f\u0120\3\2\2\2\u0120\u0122"+
		"\3\2\2\2\u0121\u011f\3\2\2\2\u0122\u0123\b.\2\2\u0123\\\3\2\2\2\u0124"+
		"\u0125\7^\2\2\u0125\u0126\t\b\2\2\u0126^\3\2\2\2\13\2\u00b4\u00ba\u00bc"+
		"\u00c2\u00cb\u0107\u0111\u011f\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
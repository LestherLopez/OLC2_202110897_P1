// Generated from c:\Users\50246\Desktop\Ingeniería-Sexto Semestre\Compiladores 2\Proyecto 1\Server\Lex.g4 by ANTLR 4.9.2
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
		CASE=18, DEFAULT=19, GUARD=20, CONTINUE=21, RETURN=22, BREAK=23, NUMBER=24, 
		STRING=25, ID=26, DIF=27, IG_IG=28, NOT=29, OR=30, AND=31, IG=32, IG_ADD=33, 
		IG_SUB=34, MAY_IG=35, MEN_IG=36, MAYOR=37, MENOR=38, MUL=39, DIV=40, ADD=41, 
		SUB=42, PERCENT=43, PARIZQ=44, PARDER=45, LLAVEIZQ=46, LLAVEDER=47, DOUBLEPTS=48, 
		QUESTION=49, POINT=50, WHITESPACE=51, COMMENT=52, LINE_COMMENT=53;
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
			"CASE", "DEFAULT", "GUARD", "CONTINUE", "RETURN", "BREAK", "NUMBER", 
			"STRING", "ID", "DIF", "IG_IG", "NOT", "OR", "AND", "IG", "IG_ADD", "IG_SUB", 
			"MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", "DIV", "ADD", "SUB", "PERCENT", 
			"PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "DOUBLEPTS", "QUESTION", 
			"POINT", "WHITESPACE", "COMMENT", "LINE_COMMENT", "ESC_SEQ"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'String'", "'Int'", "'Float'", "'Bool'", "'Character'", "'var'", 
			"'nil'", "'true'", "'false'", "'print'", "'if'", "'else'", "'while'", 
			"'let'", "'for'", "'in'", "'switch'", "'case'", "'default'", "'guard'", 
			"'continue'", "'return'", "'break'", null, null, null, "'!='", "'=='", 
			"'!'", "'||'", "'&&'", "'='", "'+='", "'-='", "'>='", "'<='", "'>'", 
			"'<'", "'*'", "'/'", "'+'", "'-'", "'%'", "'('", "')'", "'{'", "'}'", 
			"':'", "'?'", "'.'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "STRINGS", "INTS", "FLOATS", "BOOLS", "CHARACTERS", "VAR", "NIL", 
			"TRUE", "FALSE", "PRINT", "IF", "ELSE", "WHILE", "LET", "FOR", "IN", 
			"SWITCH", "CASE", "DEFAULT", "GUARD", "CONTINUE", "RETURN", "BREAK", 
			"NUMBER", "STRING", "ID", "DIF", "IG_IG", "NOT", "OR", "AND", "IG", "IG_ADD", 
			"IG_SUB", "MAY_IG", "MEN_IG", "MAYOR", "MENOR", "MUL", "DIV", "ADD", 
			"SUB", "PERCENT", "PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "DOUBLEPTS", 
			"QUESTION", "POINT", "WHITESPACE", "COMMENT", "LINE_COMMENT"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\67\u0169\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31"+
		"\t\31\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t"+
		" \4!\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t"+
		"+\4,\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64"+
		"\t\64\4\65\t\65\4\66\t\66\4\67\t\67\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\3\3"+
		"\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6"+
		"\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3"+
		"\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3"+
		"\f\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3"+
		"\17\3\20\3\20\3\20\3\20\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\22\3"+
		"\22\3\23\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3"+
		"\25\3\25\3\25\3\25\3\25\3\25\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3"+
		"\26\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\30\3\30\3\30\3\30\3\30\3\30\3"+
		"\31\6\31\u00f3\n\31\r\31\16\31\u00f4\3\31\3\31\6\31\u00f9\n\31\r\31\16"+
		"\31\u00fa\5\31\u00fd\n\31\3\32\3\32\7\32\u0101\n\32\f\32\16\32\u0104\13"+
		"\32\3\32\3\32\3\33\3\33\7\33\u010a\n\33\f\33\16\33\u010d\13\33\3\34\3"+
		"\34\3\34\3\35\3\35\3\35\3\36\3\36\3\37\3\37\3\37\3 \3 \3 \3!\3!\3\"\3"+
		"\"\3\"\3#\3#\3#\3$\3$\3$\3%\3%\3%\3&\3&\3\'\3\'\3(\3(\3)\3)\3*\3*\3+\3"+
		"+\3,\3,\3-\3-\3.\3.\3/\3/\3\60\3\60\3\61\3\61\3\62\3\62\3\63\3\63\3\64"+
		"\6\64\u0148\n\64\r\64\16\64\u0149\3\64\3\64\3\65\3\65\3\65\3\65\7\65\u0152"+
		"\n\65\f\65\16\65\u0155\13\65\3\65\3\65\3\65\3\65\3\65\3\66\3\66\3\66\3"+
		"\66\7\66\u0160\n\66\f\66\16\66\u0163\13\66\3\66\3\66\3\67\3\67\3\67\3"+
		"\u0153\28\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33"+
		"\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67"+
		"\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62c\63e\64g\65"+
		"i\66k\67m\2\3\2\t\3\2\62;\3\2$$\4\2C\\c|\6\2\62;C\\aac|\6\2\13\f\17\17"+
		"\"\"^^\4\2\f\f\17\17\t\2\"#%%--/\60<<BB]_\2\u016f\2\3\3\2\2\2\2\5\3\2"+
		"\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21"+
		"\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2"+
		"\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3"+
		"\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3"+
		"\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3"+
		"\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2"+
		"\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2"+
		"Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3"+
		"\2\2\2\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2\3o\3\2\2\2\5v\3\2\2\2\7z\3\2\2"+
		"\2\t\u0080\3\2\2\2\13\u0085\3\2\2\2\r\u008f\3\2\2\2\17\u0093\3\2\2\2\21"+
		"\u0097\3\2\2\2\23\u009c\3\2\2\2\25\u00a2\3\2\2\2\27\u00a8\3\2\2\2\31\u00ab"+
		"\3\2\2\2\33\u00b0\3\2\2\2\35\u00b6\3\2\2\2\37\u00ba\3\2\2\2!\u00be\3\2"+
		"\2\2#\u00c1\3\2\2\2%\u00c8\3\2\2\2\'\u00cd\3\2\2\2)\u00d5\3\2\2\2+\u00db"+
		"\3\2\2\2-\u00e4\3\2\2\2/\u00eb\3\2\2\2\61\u00f2\3\2\2\2\63\u00fe\3\2\2"+
		"\2\65\u0107\3\2\2\2\67\u010e\3\2\2\29\u0111\3\2\2\2;\u0114\3\2\2\2=\u0116"+
		"\3\2\2\2?\u0119\3\2\2\2A\u011c\3\2\2\2C\u011e\3\2\2\2E\u0121\3\2\2\2G"+
		"\u0124\3\2\2\2I\u0127\3\2\2\2K\u012a\3\2\2\2M\u012c\3\2\2\2O\u012e\3\2"+
		"\2\2Q\u0130\3\2\2\2S\u0132\3\2\2\2U\u0134\3\2\2\2W\u0136\3\2\2\2Y\u0138"+
		"\3\2\2\2[\u013a\3\2\2\2]\u013c\3\2\2\2_\u013e\3\2\2\2a\u0140\3\2\2\2c"+
		"\u0142\3\2\2\2e\u0144\3\2\2\2g\u0147\3\2\2\2i\u014d\3\2\2\2k\u015b\3\2"+
		"\2\2m\u0166\3\2\2\2op\7U\2\2pq\7v\2\2qr\7t\2\2rs\7k\2\2st\7p\2\2tu\7i"+
		"\2\2u\4\3\2\2\2vw\7K\2\2wx\7p\2\2xy\7v\2\2y\6\3\2\2\2z{\7H\2\2{|\7n\2"+
		"\2|}\7q\2\2}~\7c\2\2~\177\7v\2\2\177\b\3\2\2\2\u0080\u0081\7D\2\2\u0081"+
		"\u0082\7q\2\2\u0082\u0083\7q\2\2\u0083\u0084\7n\2\2\u0084\n\3\2\2\2\u0085"+
		"\u0086\7E\2\2\u0086\u0087\7j\2\2\u0087\u0088\7c\2\2\u0088\u0089\7t\2\2"+
		"\u0089\u008a\7c\2\2\u008a\u008b\7e\2\2\u008b\u008c\7v\2\2\u008c\u008d"+
		"\7g\2\2\u008d\u008e\7t\2\2\u008e\f\3\2\2\2\u008f\u0090\7x\2\2\u0090\u0091"+
		"\7c\2\2\u0091\u0092\7t\2\2\u0092\16\3\2\2\2\u0093\u0094\7p\2\2\u0094\u0095"+
		"\7k\2\2\u0095\u0096\7n\2\2\u0096\20\3\2\2\2\u0097\u0098\7v\2\2\u0098\u0099"+
		"\7t\2\2\u0099\u009a\7w\2\2\u009a\u009b\7g\2\2\u009b\22\3\2\2\2\u009c\u009d"+
		"\7h\2\2\u009d\u009e\7c\2\2\u009e\u009f\7n\2\2\u009f\u00a0\7u\2\2\u00a0"+
		"\u00a1\7g\2\2\u00a1\24\3\2\2\2\u00a2\u00a3\7r\2\2\u00a3\u00a4\7t\2\2\u00a4"+
		"\u00a5\7k\2\2\u00a5\u00a6\7p\2\2\u00a6\u00a7\7v\2\2\u00a7\26\3\2\2\2\u00a8"+
		"\u00a9\7k\2\2\u00a9\u00aa\7h\2\2\u00aa\30\3\2\2\2\u00ab\u00ac\7g\2\2\u00ac"+
		"\u00ad\7n\2\2\u00ad\u00ae\7u\2\2\u00ae\u00af\7g\2\2\u00af\32\3\2\2\2\u00b0"+
		"\u00b1\7y\2\2\u00b1\u00b2\7j\2\2\u00b2\u00b3\7k\2\2\u00b3\u00b4\7n\2\2"+
		"\u00b4\u00b5\7g\2\2\u00b5\34\3\2\2\2\u00b6\u00b7\7n\2\2\u00b7\u00b8\7"+
		"g\2\2\u00b8\u00b9\7v\2\2\u00b9\36\3\2\2\2\u00ba\u00bb\7h\2\2\u00bb\u00bc"+
		"\7q\2\2\u00bc\u00bd\7t\2\2\u00bd \3\2\2\2\u00be\u00bf\7k\2\2\u00bf\u00c0"+
		"\7p\2\2\u00c0\"\3\2\2\2\u00c1\u00c2\7u\2\2\u00c2\u00c3\7y\2\2\u00c3\u00c4"+
		"\7k\2\2\u00c4\u00c5\7v\2\2\u00c5\u00c6\7e\2\2\u00c6\u00c7\7j\2\2\u00c7"+
		"$\3\2\2\2\u00c8\u00c9\7e\2\2\u00c9\u00ca\7c\2\2\u00ca\u00cb\7u\2\2\u00cb"+
		"\u00cc\7g\2\2\u00cc&\3\2\2\2\u00cd\u00ce\7f\2\2\u00ce\u00cf\7g\2\2\u00cf"+
		"\u00d0\7h\2\2\u00d0\u00d1\7c\2\2\u00d1\u00d2\7w\2\2\u00d2\u00d3\7n\2\2"+
		"\u00d3\u00d4\7v\2\2\u00d4(\3\2\2\2\u00d5\u00d6\7i\2\2\u00d6\u00d7\7w\2"+
		"\2\u00d7\u00d8\7c\2\2\u00d8\u00d9\7t\2\2\u00d9\u00da\7f\2\2\u00da*\3\2"+
		"\2\2\u00db\u00dc\7e\2\2\u00dc\u00dd\7q\2\2\u00dd\u00de\7p\2\2\u00de\u00df"+
		"\7v\2\2\u00df\u00e0\7k\2\2\u00e0\u00e1\7p\2\2\u00e1\u00e2\7w\2\2\u00e2"+
		"\u00e3\7g\2\2\u00e3,\3\2\2\2\u00e4\u00e5\7t\2\2\u00e5\u00e6\7g\2\2\u00e6"+
		"\u00e7\7v\2\2\u00e7\u00e8\7w\2\2\u00e8\u00e9\7t\2\2\u00e9\u00ea\7p\2\2"+
		"\u00ea.\3\2\2\2\u00eb\u00ec\7d\2\2\u00ec\u00ed\7t\2\2\u00ed\u00ee\7g\2"+
		"\2\u00ee\u00ef\7c\2\2\u00ef\u00f0\7m\2\2\u00f0\60\3\2\2\2\u00f1\u00f3"+
		"\t\2\2\2\u00f2\u00f1\3\2\2\2\u00f3\u00f4\3\2\2\2\u00f4\u00f2\3\2\2\2\u00f4"+
		"\u00f5\3\2\2\2\u00f5\u00fc\3\2\2\2\u00f6\u00f8\7\60\2\2\u00f7\u00f9\t"+
		"\2\2\2\u00f8\u00f7\3\2\2\2\u00f9\u00fa\3\2\2\2\u00fa\u00f8\3\2\2\2\u00fa"+
		"\u00fb\3\2\2\2\u00fb\u00fd\3\2\2\2\u00fc\u00f6\3\2\2\2\u00fc\u00fd\3\2"+
		"\2\2\u00fd\62\3\2\2\2\u00fe\u0102\7$\2\2\u00ff\u0101\n\3\2\2\u0100\u00ff"+
		"\3\2\2\2\u0101\u0104\3\2\2\2\u0102\u0100\3\2\2\2\u0102\u0103\3\2\2\2\u0103"+
		"\u0105\3\2\2\2\u0104\u0102\3\2\2\2\u0105\u0106\7$\2\2\u0106\64\3\2\2\2"+
		"\u0107\u010b\t\4\2\2\u0108\u010a\t\5\2\2\u0109\u0108\3\2\2\2\u010a\u010d"+
		"\3\2\2\2\u010b\u0109\3\2\2\2\u010b\u010c\3\2\2\2\u010c\66\3\2\2\2\u010d"+
		"\u010b\3\2\2\2\u010e\u010f\7#\2\2\u010f\u0110\7?\2\2\u01108\3\2\2\2\u0111"+
		"\u0112\7?\2\2\u0112\u0113\7?\2\2\u0113:\3\2\2\2\u0114\u0115\7#\2\2\u0115"+
		"<\3\2\2\2\u0116\u0117\7~\2\2\u0117\u0118\7~\2\2\u0118>\3\2\2\2\u0119\u011a"+
		"\7(\2\2\u011a\u011b\7(\2\2\u011b@\3\2\2\2\u011c\u011d\7?\2\2\u011dB\3"+
		"\2\2\2\u011e\u011f\7-\2\2\u011f\u0120\7?\2\2\u0120D\3\2\2\2\u0121\u0122"+
		"\7/\2\2\u0122\u0123\7?\2\2\u0123F\3\2\2\2\u0124\u0125\7@\2\2\u0125\u0126"+
		"\7?\2\2\u0126H\3\2\2\2\u0127\u0128\7>\2\2\u0128\u0129\7?\2\2\u0129J\3"+
		"\2\2\2\u012a\u012b\7@\2\2\u012bL\3\2\2\2\u012c\u012d\7>\2\2\u012dN\3\2"+
		"\2\2\u012e\u012f\7,\2\2\u012fP\3\2\2\2\u0130\u0131\7\61\2\2\u0131R\3\2"+
		"\2\2\u0132\u0133\7-\2\2\u0133T\3\2\2\2\u0134\u0135\7/\2\2\u0135V\3\2\2"+
		"\2\u0136\u0137\7\'\2\2\u0137X\3\2\2\2\u0138\u0139\7*\2\2\u0139Z\3\2\2"+
		"\2\u013a\u013b\7+\2\2\u013b\\\3\2\2\2\u013c\u013d\7}\2\2\u013d^\3\2\2"+
		"\2\u013e\u013f\7\177\2\2\u013f`\3\2\2\2\u0140\u0141\7<\2\2\u0141b\3\2"+
		"\2\2\u0142\u0143\7A\2\2\u0143d\3\2\2\2\u0144\u0145\7\60\2\2\u0145f\3\2"+
		"\2\2\u0146\u0148\t\6\2\2\u0147\u0146\3\2\2\2\u0148\u0149\3\2\2\2\u0149"+
		"\u0147\3\2\2\2\u0149\u014a\3\2\2\2\u014a\u014b\3\2\2\2\u014b\u014c\b\64"+
		"\2\2\u014ch\3\2\2\2\u014d\u014e\7\61\2\2\u014e\u014f\7,\2\2\u014f\u0153"+
		"\3\2\2\2\u0150\u0152\13\2\2\2\u0151\u0150\3\2\2\2\u0152\u0155\3\2\2\2"+
		"\u0153\u0154\3\2\2\2\u0153\u0151\3\2\2\2\u0154\u0156\3\2\2\2\u0155\u0153"+
		"\3\2\2\2\u0156\u0157\7,\2\2\u0157\u0158\7\61\2\2\u0158\u0159\3\2\2\2\u0159"+
		"\u015a\b\65\2\2\u015aj\3\2\2\2\u015b\u015c\7\61\2\2\u015c\u015d\7\61\2"+
		"\2\u015d\u0161\3\2\2\2\u015e\u0160\n\7\2\2\u015f\u015e\3\2\2\2\u0160\u0163"+
		"\3\2\2\2\u0161\u015f\3\2\2\2\u0161\u0162\3\2\2\2\u0162\u0164\3\2\2\2\u0163"+
		"\u0161\3\2\2\2\u0164\u0165\b\66\2\2\u0165l\3\2\2\2\u0166\u0167\7^\2\2"+
		"\u0167\u0168\t\b\2\2\u0168n\3\2\2\2\13\2\u00f4\u00fa\u00fc\u0102\u010b"+
		"\u0149\u0153\u0161\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
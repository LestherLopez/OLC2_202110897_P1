import React from 'react';
import CodeMirror from '@uiw/react-codemirror';
import { javascript } from '@codemirror/lang-javascript';


function Editor(props) {
  const onChange = React.useCallback((value, viewUpdate) => {
    console.log('value:', value);
    props.input(value);
  }, []);
  return (
    <CodeMirror
    value={props.value}
    height="200px"

    extensions={[javascript({ jsx: true })]}
    onChange={onChange}
  />
  );
}
export default Editor;
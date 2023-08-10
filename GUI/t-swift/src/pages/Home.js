
import React, { useState } from "react";
import { Col, Container, Row } from "react-bootstrap";
import Button from 'react-bootstrap/Button';
import Editor from "../components/Editor";
import Consola from "../components/Consola";
import axios from "axios";
import Form from 'react-bootstrap/Form';

function Home(){
    const [editor, setEditor] = useState("");
    const [consola, setConsola] = useState("");
    const [code, setCode] = useState("print('hello world!');");
     //estados de modal para ast
     const [show, setShow] = useState(false);
     const handleClose = () => setShow(false);
     const handleShow = () => setShow(true);
     //estados de modal para tabla de simbolos
     
     const [shows, setShows] = useState(false);
     const handleCloses = () => setShows(false);
     const handleShows = () => setShows(true);
     //estados de modal para tabla de errores
     const [showe, setShowe] = useState(false);
     const handleClosee = () => setShowe(false);
     const handleShowe = () => setShowe(true);
     const handleFileChange = (event) => {
     const files = event.target.files;
     const reader = new FileReader();
     reader.onload = (e) => {
       setEditor(e.target.result);
       setCode(e.target.result)
     };
     reader.readAsText(files[0]);
   };
 
    const interpretar = async () => {
        console.log("ejecutando")
        try {
            setConsola("ejecutando...");
            if(editor ==""){
                setConsola("No hay codigo para interpretar");
                console.log("No hay codigo para interpretar");
            }else {
                console.log(editor)
                const response = await axios.post('http://localhost:5000/interpreter', {code:editor});
                console.log(response.data);
                const {consola,errores} = response.data;   
                console.log(consola);
                setConsola(response.data);
            }
        } catch (error) {
            console.log(error);
            setConsola("Error en el servidor");
        }
    };



    return (
        <Container>
            <Row>
                <Col>
                    <h1>Editor</h1>
                </Col>
                <Col>
                    <h1>Consola</h1>
                </Col>

            </Row>
            <Row>
                <Col style={{ textAlign: 'left' }}>
                    <Editor input={setEditor} value={code}/>
                </Col>
                
                <Col style={{ textAlign: 'left' }}>
                    <Consola consola={consola}/>
                </Col>
              
            </Row>
            {/*seccion de botones */}
            <Row>
                <Col>
                    <Button variant="outline-secondary" >Run</Button>{' '}
                </Col>
            </Row>
            {/*seccion de botones */}
            <Row>
                    <Col>
                        <Button variant="outline-secondary" onClick={() =>interpretar()} >Run</Button>{' '}
                        <Button variant="primary" onClick={handleShow}>
                                                                        Reporte AST
                                                                    </Button>{' '}
                        <Button variant="primary" onClick={handleShows}>Reporte Tabla de Simbolos</Button>{' '}               
                        <Button variant="primary" onClick={handleShowe}>Reporte Tabla de Errores</Button>{' '}     
                        <Form.Group controlId="formFileMultiple" className="mb-3">
            <Form.Label>Carga de Archivo TW</Form.Label>
            <Form.Control type="file" multiple onChange={handleFileChange}/>
        </Form.Group>                         
                    </Col>
                </Row>

        </Container>
    );
}

export default Home;
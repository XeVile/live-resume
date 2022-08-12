import React, {useState, useEffect} from 'react';
import logo from './logo.svg';
import './App.css';
import {Button, Container, Modal } from 'react-bootstrap';
import BasicForm from './components/Forms/BasicForm';
import EducationForm from './components/Forms/EducationForm';
import JobForm from './components/Forms/JobForm';
import SkillForm from './components/Forms/SkillForm';
import ProjectForm from './components/Forms/ProjectForm';
import ListForm from './components/Forms/ListForm';

function App() {
  const form_id = ["basic", "education", "skill", "job", "project"]

  const [userData, setUserData] = useState(false)
  const [refreshData, setRefreshData] = useState(false)

  const [changeData, setChangeData] = useState({"change": false, "id": 0})

  const [addBasic, setAddBasic] = useState(false)
  const [addEdu, setAddEdu] = useState(false)
  const [addSkill, setAddSkill] = useState(false)
  const [addJob, setAddJob] = useState(false)
  const [addProject, setAddProject] = useState(false)
  const [addList, setAddList] = useState(false)

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" height={60} width={60} />
        <h3>Live-Resume</h3>
      </header>

      <body>
        <Container>
          <Button onClick={() => setAddBasic(true)}>Add Basic Info</Button>
        </Container>


        <Modal show={addBasic} onHide={() => setAddBasic(false)}
          backdrop="static" 
          keyboard={false} 
          size="lg"
          centered>
          <Modal.Header closeButton>
            <Modal.Title>Add test</Modal.Title>
          </Modal.Header>

          <Modal.Body>
            {form_id.map(form => (
              (form === "basic") ? <div><br/><h5>Basic</h5><BasicForm /></div> :
              (form === "education") ? <div><br/><h5>Education</h5><EducationForm /></div> :
              (form === "skill") ? <div><br/><h5>Skill</h5><SkillForm /></div> :
              (form === "job") ? <div><br/><h5>Job</h5><JobForm /></div> :
              (form === "project") ? <div><br/><h5>Project</h5><ProjectForm /></div> :
              <ListForm />))
          }
          </Modal.Body>

          <Modal.Footer>
            <Button onClick={() => addBasic()}>Confirm</Button>
            <Button onClick={() => setAddBasic(false)}>Cancel</Button>
          </Modal.Footer>
        </Modal>
      </body>
    </div>
  );

  const handleSubmit = () => {
    console.log("Submit")
  }

  const handleChange = () => {
    console.log("Input recv")
  }
}

export default App;

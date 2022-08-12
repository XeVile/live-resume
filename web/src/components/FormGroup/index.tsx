import React, {useState, useEffect} from 'react';
import axios from "axios";
import {Button, Form, Container, Modal } from 'react-bootstrap';
import Section from '../sections';
import BasicForm from '../Forms/BasicForm';
import EducationForm from '../Forms/EducationForm';
import SkillForm from '../Forms/SkillForm';
import JobForm from '../Forms/JobForm';
import ProjectForm from '../Forms/ProjectForm';
import ListForm from '../Forms/ListForm';

export type Props = {
  id: string;
  title?: string;
}

const FormGroup: React.FC<Props> = (
  id: string,
  title?: string
) => {
  //
  // State Management
  //
  const [userData, setUserData] = useState(false)
  const [refreshData, setRefreshData] = useState(false)

  const [changeData, setChangeData] = useState({"change": false, "id": 0})

  const [addNew, setAddNew] = useState(false)

  useEffect(() => {
    getUserData();
  }, [])

  if (refreshData) {
    setRefreshData(false);
    getUserData();
  }

  return (
    <div>
      <Container>
        <Button onClick={() => setAddNew(true)}>Add {title} Info</Button>
      </Container>

      <Modal show={addNew} onHide={() => setAddNew(false)} centered>
        <Modal.Header closeButton>
          <Modal.Title>Add {title}</Modal.Title>
        </Modal.Header>

        <Modal.Body>
          //
          // Load the correct form
          // 
          <BasicForm />

          <Button onClick={() => addData()}>Confirm</Button>
          <Button onClick={() => setAddNew(false)}>Cancel</Button>
        </Modal.Body>
      </Modal>
    </div>
  )

  function addData() {
    setUserData(false);
    let url = "http://locahost:8880/user/create";
    axios.get(url,{})
  }

  function getUserData(user_id?: string) {
    let url = "http://locahost:8880/user/" + user_id;
    axios.get(url,
      {responseType: 'json'}
    ).then(resp => {
        if (resp.status === 200) {
          setUserData(resp.data);
        }
      })
  }

  function deleteUserData(user_id: string) {
    let url = "http://locahost:8880/user/delete/" + user_id;
    axios.delete(url, {}
    ).then(resp => {
        if (resp.status === 200) {
          setRefreshData(true)
        }
      })
  }
}

export default FormGroup;

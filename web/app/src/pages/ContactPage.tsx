import React, {FC, FormEvent, useState} from 'react';
import {useMutation} from "@apollo/client";
import {CONTACT_US} from "../graph/queries";
import {Contact, ContactVariables} from "../graph/Contact";

const ContactPage: FC = () => {
  const [name, setName] = useState('');
  const [mail, setMail] = useState('');
  const [message, setMessage] = useState('');
  const [validation, setValidation] = useState('');

  const [doSent, {data, error}] = useMutation<Contact, ContactVariables>(CONTACT_US);

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();

    if (name.length < 1 || mail.length < 1 || message.length < 2) {
      setValidation("please enter all form fields");
      return
    }

    try {
      await doSent({
        variables: {
          input: {
            email: mail,
            message: message,
            fullName: name
          }
        }
      });
    } catch (ex) {
      console.error(ex);
    }
  }


  return (
      <div className="row">
        <div className="col-md-8 offset-md-2">

          <form id="contact-form" onSubmit={handleSubmit} className="mt-md-5">
            <h3>Contact Us</h3>
            {data && <div className="alert alert-primary" role="alert">
              {data.contact}
            </div>}
            {error && <div className="alert alert-warning fade show" role="alert">
              {error.message}
            </div>}
            {validation.length > 0 && <div className="alert alert-warning fade show alert-dismissible" role="alert">
              <p className="mb-0">{validation}</p>
              <button type="button" className="close" data-dismiss="alert" aria-label="Close" onClick={() => setValidation('')}>
                <span aria-hidden="true">&times;</span>
              </button>
            </div>}
            <div className="row mt-3">
              <div className="col-md-6">
                <div className="form-group">
                  <label htmlFor="form_email">Email *</label>
                  <input id="form_email" type="email" name="email" className="form-control" placeholder="Please enter your email *" required onChange={event => setMail(event.target.value)}/>
                </div>
              </div>
              <div className="col-md-6">
                <div className="form-group">
                  <label htmlFor="form_name">Full Name *</label>
                  <input id="form_name" type="text" name="name" className="form-control" placeholder="Please enter your full name" required onChange={event => setName(event.target.value)}/>
                </div>
              </div>
            </div>
            <div className="row mt-3">
              <div className="col-md-12">
                <div className="form-group">
                  <label htmlFor="form_message">Message *</label>
                  <textarea id="form_message" name="message" className="form-control" placeholder="Message for me *" rows={4} required onChange={event => setMessage(event.target.value)}/>
                </div>
              </div>
            </div>
            <div className="row mt-3">
              <div className="col-md-12">
                <input className="btn btn-primary btn-lg" value="Send message" onClick={handleSubmit}/>
              </div>
            </div>
          </form>
        </div>
      </div>
  )
};

export default ContactPage;

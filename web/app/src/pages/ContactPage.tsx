import {createRef, FC, FormEvent, useState} from 'react';
import {gql, useMutation} from "@apollo/client";
import ReCAPTCHA from "react-google-recaptcha";
import {CAPTCHA_KEY} from "../Constants";
import {Contact, ContactVariables} from "./__generated__/Contact";


const CONTACT_US = gql`
  mutation Contact($input:ContactInput! $id:ID, $captcha:String!) {
    contact(input:$input, tweetID: $id, capthca: $captcha)
  }`;


type ContactPageProps = {
  tweetID?: string
}

const ContactPage: FC<ContactPageProps> = ({tweetID}) => {
  const [name, setName] = useState('');
  const [mail, setMail] = useState('');
  const [message, setMessage] = useState('');
  const [validation, setValidation] = useState('');
  const recaptchaRef = createRef<ReCAPTCHA>();

  const [doSent, {data, error}] = useMutation<Contact, ContactVariables>(CONTACT_US);

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();

    if (name.length < 1 || mail.length < 1 || message.length < 2) {
      setValidation("please enter all form fields");
      return
    }
    try {

      const token = await recaptchaRef.current?.executeAsync();

      await doSent({
        variables: {
          input: {
            email: mail,
            message: message,
            fullName: name
          },
          id: tweetID,
          captcha: token ?? ""
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
            <ReCAPTCHA
                ref={recaptchaRef}
                size="invisible"
                sitekey={CAPTCHA_KEY}
            />

            {!tweetID && <h3>Contact Us</h3>}

            {data && <div className="alert alert-primary" role="alert">
              {data.contact}
            </div>}
            {error && <div className="alert alert-warning fade show" role="alert">
              {error.message}
            </div>}
            {validation.length > 0 && <div className="alert alert-warning fade show alert-dismissible" role="alert">
              <p className="mb-0">{validation}</p>
              <button type="button" className="btn-close" data-dismiss="alert" aria-label="Close" onClick={() => setValidation('')}> </button>
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

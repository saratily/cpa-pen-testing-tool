import { useContext, useState, useEffect, useCallback } from "react";

import AuthContext from "../../store/auth-context";
import Errors from "../Errors/Errors";
import ToolsList from "./ToolsLists";

const Tools = (props) => {
  const authContext = useContext(AuthContext);
  const [tools, setTools] = useState([]);
  const [errors, setErrors] = useState({});
  const url = '/api/tools/' + props.uuid + '/' + props.type;

  const fetchToolsHandler = useCallback(async () => {
    setErrors({});

    try {
      const response = await fetch(url,
        {
          headers: {
            'Authorization': 'Bearer ' + authContext.token,
          },
        }
      );
      const data = await response.json();
      if (!response.ok) {
        let errorText = 'Fetching tools failed.';
        if (!data.hasOwnProperty('error')) {
          throw new Error(errorText);
        }
        if ((typeof data['error'] === 'string')) {
          setErrors({ 'unknown': data['error'] })
        } else {
          setErrors(data['error']);
        }
      } else {
        setTools(data.data);
      }
    } catch (error) {
      setErrors({ "error": error.message });
    }
  }, [authContext.token]);

  useEffect(() => {
    fetchToolsHandler();
  }, [fetchToolsHandler]);

  const addToolHandler = (toolData) => {
    setTools((prevState) => { return [...prevState, toolData] });
  }

  const deleteToolHandler = (toolID) => {
    setTools((prevState) => {
      return prevState.filter(tool => { return tool.ID !== toolID; })
    })
  }

  const editToolHandler = () => {
    fetchToolsHandler();
  }

  const toolsContent = tools.length === 0 ?
    <p>No Tools yet</p>
    :
    <ToolsList
      uuid={props.uuid}
      type={props.type}
      tools={tools}
      onEditTool={editToolHandler}
      onDeleteTool={deleteToolHandler} />;

  const errorContent = Object.keys(errors).length === 0 ? null : Errors(errors);

  return (
    <section>
      <h3 className="pb-4">{props.title}</h3>
      {errorContent}
      {toolsContent}
    </section>
  );
};

export default Tools;

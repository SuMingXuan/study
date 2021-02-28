
function handelButtonClick(){
  alert(123)
}

function TodoList() {

  return (
    <div className="TodoList">
      <div>
        <button onClick={handelButtonClick}>add</button>
      </div>
      <div>
        <ul>
          <li>ruby</li>
          <li>react</li>
          <li>go</li>
        </ul>
      </div>
    </div>
  );
}

export default TodoList;
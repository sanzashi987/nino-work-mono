## Usage
Implement the abstract class `TailEditor` and wrapped by the context `TailEditorContext`
``` tsx
  class FlowChart extends TailEditor{
    // Implement abstract methods like `onDrop` etc..
    onDrop(){}


    // use `super.render()` to extend the returned React JSX
    render(){
      return <div>
        <Toolbox>...</Toolbox>
        {super.render()}
        <Dock>...</Dock>
      </div>

    }
  }


const Editor = ()=>{
  const props = {} // collection the props according to the type 
  return (
    <MenuContext.Provider value={callbacks}>
      <FlowChart {...props}/>
    </MenuContex.Provider>
  )
}


```
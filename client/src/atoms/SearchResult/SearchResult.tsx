import { memo } from "react"

const SearchResult = memo(function SearchResult({ keyword, total }: { keyword: string, total: number }){
    console.log("search result rendered", keyword, total)

    if(keyword === ""){
        return
    }

    if(total > 0){
        return <div>{ total } Results for: <i>{ keyword }</i></div>
    }else{
        return <div>No results</div>
    }
})

export default SearchResult
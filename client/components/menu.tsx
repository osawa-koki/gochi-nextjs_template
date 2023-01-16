import Link from "next/link";
import { useEffect, useState } from "react";
import pages from "../pages";

function Menu() {

  const [current_page, setCurrentPage] = useState<string | null>(null);

  useEffect(() => {
    setCurrentPage(window.location.pathname);
    console.log(current_page);

  }, []);

  const PageChanged = () => {
    setCurrentPage(window.location.pathname);
    console.log(current_page);
  };

  return (
    <div id="Menu">
    {
      pages.map((page, index: number) => {
        return (
          <Link key={index} href={page.path} className={`btn ${current_page === page.path ? 'btn-primary' : ''}`} onClick={PageChanged}>
            {page.name}
          </Link>
        )
      })
    }
    </div>
  );
};

export default Menu;

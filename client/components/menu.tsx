import Link from "next/link";
import pages from "../pages";

function Menu(current_page: string) {
  return (
    <div id="Menu">
    {
      pages.map((page, index: number) => {
        return (
          <Link key={index} href={page.path} className={`btn ${current_page === page.path ? 'btn-primary' : ''}`}>
            {page.name}
          </Link>
        )
      })
    }
    </div>
  );
};

export default Menu;

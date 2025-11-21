import {
  BookOpenIcon,
  BooksIcon,
  BuildingOfficeIcon,
  HouseIcon,
  PersonIcon,
  UserListIcon,
} from "@phosphor-icons/react/dist/ssr";
import Link from "next/link";

const Sidebar = ({ children }: Readonly<{ children: React.ReactNode }>) => {
  return (
    <>
      <div className="drawer drawer-open">
        <input id="my-drawer-4" type="checkbox" className="drawer-toggle" />
        <div className="drawer-content">{children}</div>

        <div className="drawer-side is-drawer-close:overflow-visible ">
          <label
            htmlFor="my-drawer-4"
            aria-label="close sidebar"
            className="drawer-overlay"
          ></label>
          <div className="is-drawer-close:w-14 is-drawer-open:w-64 bg-base-200 flex flex-col items-start min-h-full">
            <div className="flex items-center text-center justify-center w-full">
              <div className="flex flex-col py-6 is-drawer-close:py-8 justify-center items-center">
                <BooksIcon weight="duotone" className="mb-2" size={40} />
                <h1 className="is-drawer-close:hidden font-semibold text-sm">
                  BOOK <br /> MANAGEMENT
                  <br /> APPS
                </h1>
              </div>
            </div>

            <div className="divider"></div>

            <ul className="menu w-full grow">
              <h4 className="is-drawer-close:hidden font-semibold mb-2">
                Book Management
              </h4>
              <Link href="/">
                <li>
                  <button
                    className="is-drawer-close:tooltip is-drawer-close:tooltip-right"
                    data-tip="Dashboard"
                  >
                    <HouseIcon weight="duotone" size={20} />
                    <span className="is-drawer-close:hidden">Dashboard</span>
                  </button>
                </li>
              </Link>

              <Link href="/books">
                <li>
                  <button
                    className="is-drawer-close:tooltip is-drawer-close:tooltip-right"
                    data-tip="Books"
                  >
                    <BookOpenIcon weight="duotone" size={20} />
                    <span className="is-drawer-close:hidden">Books</span>
                  </button>
                </li>
              </Link>

              <h4 className="is-drawer-close:hidden font-semibold mb-2 mt-4">
                Reference Data
              </h4>

              <Link href="/authors">
                <li>
                  <button
                    className="is-drawer-close:tooltip is-drawer-close:tooltip-right"
                    data-tip="Authors"
                  >
                    <UserListIcon weight="duotone" size={20} />
                    <span className="is-drawer-close:hidden">Authors</span>
                  </button>
                </li>
              </Link>

              <Link href="/publishers">
                <li>
                  <button
                    className="is-drawer-close:tooltip is-drawer-close:tooltip-right"
                    data-tip="Publishers"
                  >
                    <BuildingOfficeIcon weight="duotone" size={20} />
                    <span className="is-drawer-close:hidden">Publishers</span>
                  </button>
                </li>
              </Link>
            </ul>

            <div
              className="m-2 is-drawer-close:tooltip is-drawer-close:tooltip-right"
              data-tip="Open"
            >
              <label
                htmlFor="my-drawer-4"
                className="btn btn-ghost btn-circle drawer-button is-drawer-open:rotate-y-180"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 24 24"
                  strokeLinejoin="round"
                  strokeLinecap="round"
                  strokeWidth="2"
                  fill="none"
                  stroke="currentColor"
                  className="inline-block size-4 my-1.5"
                >
                  <path d="M4 4m0 2a2 2 0 0 1 2 -2h12a2 2 0 0 1 2 2v12a2 2 0 0 1 -2 2h-12a2 2 0 0 1 -2 -2z"></path>
                  <path d="M9 4v16"></path>
                  <path d="M14 10l2 2l-2 2"></path>
                </svg>
              </label>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default Sidebar;

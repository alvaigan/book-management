"use client"

import { useAppStore } from "@/lib/hooks";
import { Books } from "@phosphor-icons/react";
import { BooksIcon } from "@phosphor-icons/react/dist/ssr";
import Link from "next/link";

const Login = () => {

  const store = useAppStore()

  return (
    <div className="min-h-screen flex items-center justify-center bg-base-200 p-4">
      <div className="flex flex-col justify-center items-center w-100 z-10">
        <div className="card w-full max-w-md bg-base-100 border-1 border-base-300">
          <div className="card-body">

            <div className="w-full">
              <div className="mb-6">
                <div className="flex flex-col items-center text-2xl justify-center gap-2 font-medium text-center">
                  <BooksIcon weight="duotone" className="me-2" size={60} />
                  <h1>
                    Book Management <br /> Apps
                  </h1>
                </div>
              </div>
            </div>

            <div>
              <div className="form-control">
                <label className="label">
                  <span className="label-text">Username</span>
                </label>
                <input
                  type="text"
                  placeholder="Enter your email"
                  className="input input-bordered w-full"
                />
              </div>

              <div className="form-control mt-4">
                <label className="label">
                  <span className="label-text">Password</span>
                </label>
                <input
                  type="password"
                  placeholder="Enter your password"
                  className="input input-bordered w-full"
                />
              </div>

              <div className="form-control mt-6">
                <button className="btn btn-primary w-full">Login</button>
              </div>
            </div>

            <p className="text-center mt-6 text-sm">
              If you don't have any accounts, please
              <Link href="/auth/register" className="link link-primary ml-1">
                Register
              </Link>
            </p>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Login;

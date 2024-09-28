import type { Metadata } from 'next';
import { Jost } from 'next/font/google';
import React from 'react';
import './globals.css';

const jost = Jost({
  subsets: ['latin'],
  preload: true,
  display: 'swap',
  variable: '--font-jost',
  weight: ['300', '400', '500', '600', '700'],
});

export const metadata: Metadata = {
  title: 'Lukas Grimm - Software Engineer',
  description: 'Lukas is an amazing Software Engineer from Berlin, Germany.',
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={`${jost.variable} dark:bg-black dark:text-white`}>
        {children}
      </body>
    </html>
  );
}

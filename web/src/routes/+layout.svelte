


<nav>
  <h1><a href="/">Title</a></h1>
  <ul>
    <li>Login</li>
  </ul>
</nav>

<div>
  <slot />
</div>


<style>
  :global(*) {
    box-sizing: border-box;

    --color-aquamarine: #B5FFE1;
    --color-celadon: #93E5AB;
    --color-mint: #65B891;
    --color-dark-cyan: #4E878C;
    --color-dark-green: #00241B;

    --color-gunmetal: #272D2D;
    --color-rose-quartz: #A39BA8;
    --color-powder-blue: #B8C5D6;
    --color-alice-blue: #EDF5FC;
    --color-emerald: #23CE6B;
  }

  :global(body) {
    margin: 0;
    font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
    background-color: var(--color-alice-blue);
    color: var(--color-gunmetal);
  }

  nav {
    padding: 0.25rem 0.5rem;
    width: 100%;
    background-color: var(--color-powder-blue);
    margin: 0;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
  }

  ul {
    display: flex;
    flex-direction: row;
    margin: 0;
    list-style-type: none;
    padding-left: 0;
  }

  li {
    margin-left: 0;
  }

  h1 {
    margin: 0;
    font-size: 2.5rem;
  }

  a {
    text-decoration: none;
    color: var(--color-gunmetal);
  }

  div {
    padding: 0.5rem;
  }

</style>
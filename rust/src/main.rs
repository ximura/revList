mod node;

fn create_tree() -> node::Tree {
    let mut tree = node::Tree::new();
    tree.add("Cargo.lock".to_string());
    tree.add("Cargo.toml".to_string());
    tree.add("LICENSE".to_string());
    tree.add("README.md".to_string());
    let l1 = tree.add("src".to_string());
    let l2 = l1.add("test".to_string());
    l2.add("tree_test.rs".to_string());
    l1.add("main.rs".to_string());
    l1.add("node.rs".to_string());
    let l2 = l1.add("test".to_string());
    l2.add("tree_test.rs".to_string());
    let l1 = tree.add("target".to_string());
    l1.add("CACHEDIR.TAG".to_string());
    let l2 = l1.add("debug".to_string());
    l2.add("build".to_string());
    let l3 = l2.add("deps".to_string());
    l3.add("rstree-7606b252a0b14877".to_string());
    l3.add("rstree-7606b252a0b14877.d".to_string());
    l2.add("examples".to_string());
    let l3 = l2.add("incremental".to_string());
    let l4 = l3.add("rstree-305oidwy6y7vi".to_string());
    l4.add("s-gta5zehn80-rh0smg-9tlsfp5vm7pjiz25z60vnx81s".to_string());
    l4.add("s-gta5zehn80-rh0smg.lock".to_string());

    tree
}

fn main() {
    let tree = create_tree();
    tree.println()
}
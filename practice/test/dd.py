import matplotlib.pyplot as plt

# Create figure and axis
fig, ax = plt.subplots()

# Define nodes and their positions
main_nodes = ['Market Structure Analysis', 'Behavior Analysis', 'Competition Policy Analysis', 'Platform Economy Analysis']
main_node_positions = [(1, 4), (1, 3), (1, 2), (1, 1)]

sub_nodes = [
    ['Market Share', 'Network Effects', 'Economies of Scale', 'Entry Barriers'],
    ['Consumer Cognitive Bias', 'Information Asymmetry', 'Consumer Rights', 'Merchant Rights'],
    ['Market Competition', 'Consumer Welfare', 'Legal Regulatory Challenges'],
    ['Business Model', 'Market Operation Mechanism', 'Internal Governance', 'Technical Support']
]

sub_node_positions = [
    [(2, 4.5), (2, 4), (2, 3.5), (2, 3)],
    [(2, 3.5), (2, 3), (2, 2.5), (2, 2)],
    [(2, 2.5), (2, 2), (2, 1.5)],
    [(2, 1.5), (2, 1), (2, 0.5), (2, 0)]
]

# Draw main nodes
for node, position in zip(main_nodes, main_node_positions):
    ax.text(position[0], position[1], node, fontsize=12, bbox=dict(facecolor='orange', edgecolor='black', boxstyle='round,pad=0'))